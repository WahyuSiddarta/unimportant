package structure

import (
	"database/sql"
	"encoding/json"

	"github.com/ruangnyaman/rna-ecommerce-backend/helper"
)

type SRegency struct {
	RegencyID    string         `db:"regency_id"`
	RegencyName  sql.NullString `db:"regency_name"`
	ProvinceID   sql.NullString `db:"province_id"`
	ProvinceName sql.NullString `db:"province_name"`
	CountryCode  sql.NullString `db:"country_code"`
}

func (a SRegency) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RegencyID    string `json:"regency_id"`
		RegencyName  string `json:"regency_name"`
		ProvinceID   string `json:"province_id"`
		ProvinceName string `json:"province_name"`
		CountryCode  string `json:"country_code"`
	}{
		RegencyID:    a.RegencyID,
		RegencyName:  helper.NullStringToString(a.RegencyName),
		ProvinceID:   helper.NullStringToString(a.ProvinceID),
		ProvinceName: helper.NullStringToString(a.ProvinceName),
		CountryCode:  helper.NullStringToString(a.CountryCode),
	})
}

type SProvince struct {
	ProvinceID   string `json:"province_id" db:"province_id"`
	ProvinceName string `json:"province_name" db:"province_name"`
}

//MarshalJSON : overload convert SProvince struct
func (a SProvince) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ProvinceID   string `json:"province_id" db:"province_id"`
		ProvinceName string `json:"province_name" db:"province_name"`
	}{
		ProvinceID:   a.ProvinceID,
		ProvinceName: a.ProvinceName,
	})
}

type SCountry struct {
	ID        string `json:"id" db:"id"`
	NiceName  string `json:"nicename" db:"nicename"`
	ISO3      string `json:"iso3" db:"iso3"`
	PhoneCode int    `json:"phonecode" db:"phonecode"`
}

//MarshalJSON : overload convert SCountry struct
func (a SCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        string `json:"id" db:"id"`
		NiceName  string `json:"nicename" db:"nicename"`
		ISO3      string `json:"iso3" db:"iso3"`
		PhoneCode int    `json:"phonecode" db:"phonecode"`
	}{
		ID:        a.ID,
		NiceName:  a.NiceName,
		ISO3:      a.ISO3,
		PhoneCode: a.PhoneCode,
	})
}
