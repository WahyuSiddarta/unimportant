package api

import (
	mqtt "github.com/budisugianto/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

//API :
type API struct {
	Taxes       int
	BniSecret   string
	BniClientID string
	BniPrefixVA string
	TrxURL      string
	Log         *zerolog.Logger
	DBSQL       []*sqlx.DB
	Router      *echo.Echo
	Config      *viper.Viper
	MQPub       mqtt.Client
	MQSub       mqtt.Client
}
