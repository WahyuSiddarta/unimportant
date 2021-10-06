package structure

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/WahyuSiddarta/unimportant/helper"
)

type AdminLogin struct {
}
type UserLogin struct {
	Email     sql.NullString `db:"email"`
	Password  string         `db:"password"`
	UserId    string         `db:"user_id"`
	CreatedAt time.Time      `db:"created_at"`
	Ktp       sql.NullString `db:"ktp_path"`
	Status    int            `db:"status"`
}

func (a UserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		UserId    string `json:"user_id"`
		CreatedAt string `json:"created_at"`
		Ktp       string `json:"ktp_path"`
		Status    int    `json:"status"`
	}{
		Email:     helper.NullStringToString(a.Email),
		Password:  a.Password,
		UserId:    a.UserId,
		CreatedAt: helper.TimeToString(a.CreatedAt),
		Ktp:       helper.NullStringToString(a.Ktp),
		Status:    a.Status,
	})
}
