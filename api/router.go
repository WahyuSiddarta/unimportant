package api

import (
	//"bytes"

	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"runtime"

	//"net/mail"
	//"html/template"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/budisugianto/tollbooth_echo"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/ruangnyaman/rna-ecommerce-backend/conf"
	database "github.com/ruangnyaman/rna-ecommerce-backend/db"
	"github.com/ruangnyaman/rna-ecommerce-backend/helper"
	"github.com/ruangnyaman/rna-ecommerce-backend/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	eMw "github.com/labstack/echo/v4/middleware"
	"github.com/mattn/go-colorable"
	//echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	//	"bytes"
	//	"html/template"
	//	"net/mail"
)

// LogConfig : Configuration for logging
type LogConfig struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJSON makes the log framework log JSON
	EncodeLogsAsJSON bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

func newRollingFile(config LogConfig) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}
}

var (
	// Logger :
	Logger *zerolog.Logger
	//Apix :
	Apix API
	// Version :
	Version = "1.00.01"
	// Maintenance :
	Maintenance  = false
	ServerStatus string
)

const tstamp = "02/01 15:04:05.999"

// @Schemes http https
func Init() *API {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	var writers []io.Writer
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		if runtime.GOOS == "windows" {
			writers = append(writers, zerolog.ConsoleWriter{
				Out:        colorable.NewColorableStdout(),
				TimeFormat: tstamp,
			})
		} else {
			writers = append(writers, zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: tstamp,
				FormatTimestamp: func(i interface{}) string {
					parse, _ := time.Parse(time.RFC3339Nano, i.(string))
					x, _ := helper.TimeInWIB(parse)
					return "\033[1;36m" + x.Format(tstamp) + "\033[0m"
				},
			})
		}
	}

	config := LogConfig{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJSON:      false,
		FileLoggingEnabled:    true,
		Directory:             "./log",
		Filename:              "backend.log",
		MaxSize:               100,
		MaxBackups:            7,
		MaxAge:                30,
	}

	writers = append(writers, newRollingFile(config))
	mw := io.MultiWriter(writers...)

	wr := diode.NewWriter(mw, 500, 50*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})

	Logger := zerolog.New(wr).With().Timestamp().Logger()
	Logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJSON).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("Logging system configured")

	var err error
	Apix.Config, err = conf.GetConfig()
	if err != nil {
		panic("Error opening config file")
	}
	helper.MailInit(Apix.Config)

	loglevel, err := zerolog.ParseLevel(Apix.Config.GetString("loglevel"))
	if err != nil {
		loglevel, _ = zerolog.ParseLevel("info")
	}
	Apix.Log = &Logger
	helper.Logger = &Logger
	database.Logger = &Logger
	model.Logger = &Logger

	Logger.Level(loglevel)

	//fmt.Printf("\n%+v\n", Apix.Config.AllSettings())
	Logger.Info().Msg("Router init started")
	router := echo.New()

	ServerStatus = "production"
	if ServerStatus != "production" {
		f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			panic("Error opening log file")
		}

		DefaultLoggerConfig := eMw.LoggerConfig{
			Skipper: eMw.DefaultSkipper,
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
			Output: f,
		}
		router.Use(eMw.LoggerWithConfig(DefaultLoggerConfig))

		//echopprof.Wrap(router)
	}

	// Middleware
	router.Use(eMw.SecureWithConfig(eMw.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSMaxAge:         31536000,
		//ContentSecurityPolicy: "default-src 'self'",
	}))
	router.Use(eMw.BodyLimit(Apix.Config.GetString("bodylimit")))
	router.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	router.Use(eMw.Recover())

	router.Use(eMw.CORSWithConfig(eMw.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodOptions, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderXCSRFToken, echo.HeaderAccept, "Authentication", echo.HeaderAuthorization, echo.HeaderXRealIP},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	Apix.DBSQL = [](*sqlx.DB){
		database.PSQLDBInit(),
	}

	// Config, err := conf.GetConfig()
	// if err != nil {
	// 	panic("Error opening config file")
	// }

	Logger.Info().Msg("Database Init completed")
	Logger.Info().Msgf("Router init completed")

	Apix.Router = router
	return &Apix
}

//SetupRoutes : set router
func (a *API) SetupRoutes() {

	defer func() {
		if err := recover(); err != nil {
			Logger.Error().Msgf("panic occurred: %#v\n", err)
		}
	}()

	config := eMw.JWTConfig{
		Claims:        &helper.JWTCustomClaims{},
		SigningMethod: "HS512",
		SigningKey:    []byte(a.Config.GetString("jwtsecret")),
	}

	limiter := tollbooth.NewLimiter(Apix.Config.GetFloat64("limitrate"), &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP", "CF-Connecting-IP"})

	a.Router.GET("/", func(c echo.Context) error { return c.HTML(http.StatusOK, "OK") })

	// Routes
	rpub := a.Router.Group("/public")
	rpub.Use(tollbooth_echo.LimitHandler(limiter))
	pv1 := rpub.Group("/v1")

	pv1.POST("/register", a.CreateNewUserByMail)
	pv1.POST("/login", a.Login)
	pv1.POST("/test", a.TestLogin)

	rapi := a.Router.Group("/api")
	av1 := rapi.Group("/v1") // - api/v1
	av1.Use(tollbooth_echo.LimitHandler(limiter))
	av1.Use(eMw.JWTWithConfig(config))
	av1.GET("/profile", a.GetUSerProfile)

	av1.POST("/kyc", a.UploadUserIdentity)

	go func() {
		a.Router.Logger.Info(a.Router.Start(":1323"))
	}()

	/*** graceful shutdown ***/
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Application gracefully shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Println("Application fully shutdown")
	if err := a.Router.Shutdown(ctx); err != nil {
		a.Router.Logger.Fatal(err)
	}
}
