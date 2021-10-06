package structure

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/WahyuSiddarta/unimportant/helper"
)

type PatnerHotel struct {
	OwnerID       int64          `db:"owner_id"`
	Name          string         `db:"owner_name"`
	Email         string         `db:"email"`
	PhoneNumber   string         `db:"phone_number"`
	Address       string         `db:"address"`
	RegencyID     string         `db:"regency_id"`
	ProvinceName  string         `db:"province_name"`
	PostalCode    string         `db:"postal_code"`
	CPName        string         `db:"cp_name"`
	CPPosition    string         `db:"cp_position"`
	CPEmail       string         `db:"cp_email"`
	CPPhoneNumber string         `db:"cp_phone_number"`
	CreatedAt     time.Time      `db:"created_at"`
	PropertyType  sql.NullString `db:"property_type"`
	TotalRoom     sql.NullInt64  `db:"total_room"`
}

func (a PatnerHotel) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OwnerID       int64  `json:"owner_id"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		PhoneNumber   string `json:"phone_number"`
		Address       string `json:"address"`
		RegencyID     string `json:"regency_id"`
		ProvinceName  string `json:"province_name"`
		PostalCode    string `json:"postal_code"`
		CPName        string `json:"cp_name"`
		CPPosition    string `json:"cp_position"`
		CPEmail       string `json:"cp_email"`
		CPPhoneNumber string `json:"cp_phone_number"`
		CreatedAt     string `json:"created_at"`
		PropertyType  string `json:"property_type"`
		TotalRoom     int64  `json:"total_room"`
	}{
		OwnerID:       a.OwnerID,
		Name:          a.Name,
		Email:         a.Email,
		PhoneNumber:   a.PhoneNumber,
		Address:       a.Address,
		RegencyID:     a.RegencyID,
		ProvinceName:  a.ProvinceName,
		PostalCode:    a.PostalCode,
		CPName:        a.CPName,
		CPPosition:    a.CPPosition,
		CPEmail:       a.CPEmail,
		CPPhoneNumber: a.CPPhoneNumber,
		CreatedAt:     helper.TimeToString(a.CreatedAt),
		PropertyType:  helper.NullStringToString(a.PropertyType),
		TotalRoom:     helper.NullInt64ToInt64(a.TotalRoom),
	})
}
