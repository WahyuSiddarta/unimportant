package structure

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/WahyuSiddarta/unimportant/helper"
)

type LostRecovery struct {
	UserID       string         `db:"user_id"`
	LoginID      string         `db:"login_id"`
	CreatedAt    time.Time      `db:"created_at"`
	Code1        sql.NullString `db:"code1"`
	Code2        sql.NullString `db:"code2"`
	Code3        sql.NullString `db:"code3"`
	Item1        sql.NullString `db:"item1"`
	Item2        sql.NullString `db:"item2"`
	Item3        sql.NullString `db:"item3"`
	RecoveryType int            `db:"recovery_type"`
}

func (a LostRecovery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		LoginID      string `json:"login_id"`
		CreatedAt    string `json:"created_at"`
		Code1        string `json:"code1"`
		Code2        string `json:"code2"`
		Code3        string `json:"code3"`
		Item1        string `json:"item1"`
		Item2        string `json:"item2"`
		Item3        string `json:"item3"`
		RecoveryType int    `json:"recovery_type"`
	}{
		LoginID:      a.LoginID,
		CreatedAt:    helper.TimeToString(a.CreatedAt),
		Code1:        helper.NullStringToString(a.Code1),
		Code2:        helper.NullStringToString(a.Code2),
		Code3:        helper.NullStringToString(a.Code3),
		Item1:        helper.NullStringToString(a.Item1),
		Item2:        helper.NullStringToString(a.Item2),
		Item3:        helper.NullStringToString(a.Item3),
		RecoveryType: a.RecoveryType,
	})
}
