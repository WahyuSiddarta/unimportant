package api

import (
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"strconv"

	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ruangnyaman/rna-ecommerce-backend/helper"
	"github.com/ruangnyaman/rna-ecommerce-backend/model"
	"github.com/ruangnyaman/rna-ecommerce-backend/structure"
)

func (a API) UploadUserIdentity(c echo.Context) error {
	userID := helper.GetJWTUserID(c)

	identity, err := c.FormFile("identity")
	if err != nil {
		a.Log.Error().Msgf("identity %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E9100",
			"message": "Error Upload File",
		})
	}
	fc1, err := identity.Open()
	if err != nil {
		a.Log.Error().Msgf("file error open %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code": "E9602", "message": "Cannot Open File!",
		})
	}

	bc1, err := ioutil.ReadAll(fc1)
	fc1.Close()
	if err != nil {
		a.Log.Error().Msgf("file error read %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code": "E9603", "message": "Cannot Open File!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Activation resend Succesfully",
	})
}
func (a API) GetUSerProfile(c echo.Context) error {
	userID := helper.GetJWTUserID(c)

	DBcon := model.GetDB(a.DBSQL)
	user, err := DBcon.MGetUserProfileByUserID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"code":    "E9501",
			"message": "Error!, Something Went Wrong",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":    user,
		"message": "Activation resend Succesfully",
	})
}

func (a API) Login(c echo.Context) error {
	var (
		tok  string
		err  error
		User structure.UserLogin
	)

	email := strings.ToLower(c.FormValue("email"))

	if helper.CheckEmptyString(email) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1010",
			"message": "Username / Password must be supplied",
		})
	}

	ts := c.FormValue("ts")
	if helper.CheckEmptyString(ts) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1009",
			"message": "Problem with timestamp",
		})
	}
	n, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1009",
			"message": "Problem with timestamp",
		})
	}

	bts := make([]byte, 8)
	binary.LittleEndian.PutUint64(bts, uint64(n))
	tm := helper.TimeToTimestamp(time.Now())
	if helper.ABS(tm-n) > 300 {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E2001",
			"message": "Timestamp difference too high",
		})
	}

	password := c.FormValue("password")
	if helper.CheckEmptyString(password) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1010",
			"message": "Username / Password must be supplied",
		})
	}

	DBcon := model.GetDB(a.DBSQL)
	User, err = DBcon.MGetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{
				"code":    "E1102",
				"message": "Problem with username or password",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"code":    "E9501",
			"message": "Error!, Something Went Wrong",
		})
	}

	b, _ := hex.DecodeString(User.Password)
	xb := append(b, []byte(bts)...)
	pass := helper.SHA256b(xb)
	xPass := hex.EncodeToString(pass)

	if password != xPass {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{
			"code":    "E1102",
			"message": "Problem with username or password",
		})
	}

	/* Check for verification */
	atTime, tok, err := helper.CreateJWT(email, User.UserId, "CUS", a.Config.GetString("jwtsecret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"code":    "E9502",
			"message": "Problem with Token",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login Successful",
		"user_id": User.UserId,
		"expired": atTime,
		"token":   tok,
	})
}

func (a API) CreateNewUserByMail(c echo.Context) error {
	email := strings.ToLower(c.FormValue("email"))
	if helper.CheckEmptyString(email) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1010",
			"message": "Username / Password must be supplied",
		})
	}
	password := c.FormValue("password")
	if helper.CheckEmptyString(password) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1010",
			"message": "Username / Password must be supplied",
		})
	}
	password2 := c.FormValue("password2")
	if helper.CheckEmptyString(password2) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E1010",
			"message": "Username / Password must be supplied",
		})
	}

	if password != password2 {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"code":    "E2000",
			"message": "Password Not Match",
		})
	}

	DBcon := model.GetDB(a.DBSQL)
	user_id, err := DBcon.MCreateNewUser(email, password)
	if err != nil {
		if err.Error() == "user already exist" {
			return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
				"code":    "E3000",
				"message": "User already exist",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"code":    "E9501",
			"message": "Error!, Something Went Wrong",
		})
	}

	atTime, tok, err := helper.CreateJWT(email, user_id, "CUS", a.Config.GetString("jwtsecret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"code":    "E9502",
			"message": "Problem with Token",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Register Successful",
		"expired": atTime,
		"token":   tok,
	})
}
