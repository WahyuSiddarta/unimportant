package model

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

//Debug :
var (
	Debug  = false
	Logger *zerolog.Logger
)

//const :
const (
	RECTypeEmailRecovery               = 1
	RECTypePhoneRecovery               = 2
	RECTypePasswordRecovery            = 3
	RECTypeGARecovery                  = 4
	RECTypeChangePassword              = 5
	RECTypeChangeEmailOTPCode          = 61
	RECTypeChangeEmailSMSOTPCode       = 62
	RECTypeChangeMobileOTPCode         = 7
	RECTYPEDEACTIVATE2FAGAEmailCode    = 81
	RECTYPEDEACTIVATE2FAOTHEREmailCode = 82
	RECTypeBeneficiaryOTPCode          = 10
	RECTypeWDOTP                       = 9
)

//SQLTimeFormat : Format string for golang to output SQL standar time
const SQLTimeFormat = "2006-01-02 15:04:05"

//SQLDateFormat :
const SQLDateFormat = "2006-01-02"

// TimeRange :
type TimeRange struct {
	Valid bool
	Start time.Time
	End   time.Time
}

// RetSuccess :
type RetSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//AString :
type AString []string

//JSONMap :
type JSONMap map[string]interface{}

//SQLDB : sqlx db pointer
type SQLDB struct {
	DB []*sqlx.DB
}

//GetDB : initialize DB -> pointer is transfered into model from database (localize)
func GetDB(db []*sqlx.DB) SQLDB {
	return SQLDB{
		DB: db,
	}
}

//SRowCount : DB Row count
type SRowCount struct {
	Count int64 `db:"row_count"`
}

//MarshalJSON : overload convert SMember struct
func (a SRowCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Count int64 `json:"row_count"`
	}{
		Count: a.Count,
	})
}
