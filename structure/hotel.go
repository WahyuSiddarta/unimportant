package structure

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/WahyuSiddarta/unimportant/helper"
)

type SRoomType struct {
	HotelID      string `db:"hotel_id"`
	RoomTypeID   string `db:"room_type_id"`
	RoomTypeName string `db:"room_type_name"`
}

func (a SRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID      string `json:"hotel_id"`
		RoomTypeID   string `json:"room_type_id"`
		RoomTypeName string `json:"room_type_name"`
	}{
		HotelID:      a.HotelID,
		RoomTypeID:   a.RoomTypeID,
		RoomTypeName: a.RoomTypeName,
	})
}

// hotel favorite
type SHotelFavorite struct {
	UserID       string          `db:"user_id"`
	HotelID      string          `db:"hotel_id"`
	HotelName    string          `db:"hotel_name"`
	Star         int             `db:"stars"`
	HotelType    string          `db:"hotel_type"`
	MainPitcure  string          `db:"main_picture"`
	MinCharge    int             `db:"min_charge"`
	RegencyName  string          `db:"regency_name"`
	ProvinceName string          `db:"province_name"`
	AC           int             `db:"ac"`
	Fan          int             `db:"fan"`
	Tv           int             `db:"tv"`
	Wifi         int             `db:"wifi"`
	HotCool      int             `db:"hot_cool"`
	Soap         int             `db:"soap"`
	Reception    int             `db:"reception"`
	Resto        int             `db:"resto"`
	Pool         int             `db:"pool"`
	Gym          int             `db:"gym"`
	RoomService  int             `db:"room_service"`
	Parking      int             `db:"parking"`
	Breakfast    int             `db:"breakfast"`
	Refund       int             `db:"refund"`
	GuestMax     int             `db:"guest_max"`
	RefundCharge float64         `db:"refund_charge"`
	Review       float32         `db:"review"`
	ActivePromo  sql.NullFloat64 `db:"active_promo"`
}

func (a SHotelFavorite) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID       string  `json:"user_id"`
		HotelID      string  `json:"hotel_id"`
		HotelName    string  `json:"hotel_name"`
		Star         int     `json:"stars"`
		HotelType    string  `json:"hotel_type"`
		MainPitcure  string  `json:"main_picture"`
		MinCharge    int     `json:"min_charge"`
		RegencyName  string  `json:"regency_name"`
		ProvinceName string  `json:"province_name"`
		AC           int     `json:"ac"`
		Fan          int     `json:"fan"`
		Tv           int     `json:"tv"`
		Wifi         int     `json:"wifi"`
		HotCool      int     `json:"hot_cool"`
		Soap         int     `json:"soap"`
		Reception    int     `json:"reception"`
		Resto        int     `json:"resto"`
		Pool         int     `json:"pool"`
		Gym          int     `json:"gym"`
		RoomService  int     `json:"room_service"`
		Parking      int     `json:"parking"`
		Breakfast    int     `json:"breakfast"`
		Refund       int     `json:"refund"`
		GuestMax     int     `json:"guest_max"`
		RefundCharge float64 `json:"refund_charge"`
		Review       float32 `json:"review"`
		ActivePromo  float64 `json:"active_promo"`
	}{
		UserID:       a.UserID,
		HotelID:      a.HotelID,
		HotelName:    a.HotelName,
		Star:         a.Star,
		HotelType:    a.HotelType,
		MainPitcure:  a.MainPitcure,
		MinCharge:    a.MinCharge,
		RegencyName:  a.RegencyName,
		ProvinceName: a.ProvinceName,
		AC:           a.AC,
		Fan:          a.Fan,
		Tv:           a.Tv,
		Wifi:         a.Wifi,
		HotCool:      a.HotCool,
		Soap:         a.Soap,
		Reception:    a.Reception,
		Resto:        a.Resto,
		Pool:         a.Pool,
		Gym:          a.Gym,
		RoomService:  a.RoomService,
		Parking:      a.Parking,
		Breakfast:    a.Breakfast,
		Refund:       a.Refund,
		GuestMax:     a.GuestMax,
		RefundCharge: a.RefundCharge,
		Review:       a.Review,
		ActivePromo:  helper.NullFloat64ToFloat64(a.ActivePromo),
	})
}

type SHotelsFavorite struct {
	NextPage       bool             `json:"next_page"`
	HotelsFavorite []SHotelFavorite `json:"hotels_favorite"`
}

func (a SHotelsFavorite) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage       bool             `json:"next_page"`
		HotelsFavorite []SHotelFavorite `json:"hotels_favorite"`
	}{
		NextPage:       a.NextPage,
		HotelsFavorite: a.HotelsFavorite,
	})
}

// hotels
type SHelperHotelRoom struct {
	DiscPromo    sql.NullInt64  `db:"f_promo"`
	DiscCampaign sql.NullInt64  `db:"f_campaign"`
	RoomType     sql.NullInt64  `db:"f_room_type"`
	PromoID      sql.NullString `db:"f_promo_id"`
}

type SRoom struct {
	HotelID        string `db:"hotel_id"`
	RoomType       int    `db:"room_type"`
	RoomTypeName   string `db:"room_type_name"`
	RoomOpen       bool   `db:"room_open"`
	AvailableRooms int    `db:"available_room"`
	Price          int64  `db:"price"`

	DiscPromo      sql.NullInt64 `db:"f_promo"`
	DiscCampaign   sql.NullInt64 `db:"f_campaign"`
	ValPromo       int64         `json:"promo_value"`
	ValCampaign    int64         `json:"campaign_value"`
	Taxes          int64         `json:"taxes"`
	DiscCharge     int64         `json:"min_charge_aftertax"`
	TotalMinCharge int64         `json:"total_min_charge"`

	CheckinDate  time.Time `db:"checkin_date"`
	AC           int       `db:"ac"`
	Fan          int       `db:"fan"`
	TV           int       `db:"tv"`
	Wifi         int       `db:"wifi"`
	HotCool      int       `db:"hot_cool"`
	Soap         int       `db:"soap"`
	RoomService  int       `db:"room_service"`
	Breakfast    int       `db:"breakfast"`
	Refund       int       `db:"refund"`
	GuestMax     int       `db:"guest_max"`
	RefundCharge float64   `db:"refund_charge"`
	RoomPitcure  string    `db:"room_picture"`
	RewardPoint  int64     `db:"reward_point"`
}

func (a SRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID        string `json:"hotel_id"`
		RoomType       int    `json:"room_type"`
		RoomTypeName   string `json:"room_type_name"`
		RoomOpen       bool   `json:"room_open"`
		AvailableRooms int    `json:"available_room"`
		Price          int64  `json:"price"`

		DiscPromo      int64 `json:"promo_disc"`
		DiscCampaign   int64 `json:"campaign_disc"`
		ValPromo       int64 `json:"promo_value"`
		ValCampaign    int64 `json:"campaign_value"`
		Taxes          int64 `json:"taxes"`
		DiscCharge     int64 `json:"min_charge_aftertax"`
		TotalMinCharge int64 `json:"total_min_charge"`

		CheckinDate  string  `json:"checkin_date"`
		AC           int     `json:"ac"`
		Fan          int     `json:"fan"`
		TV           int     `json:"tv"`
		Wifi         int     `json:"wifi"`
		HotCool      int     `json:"hot_cool"`
		Soap         int     `json:"soap"`
		RoomService  int     `json:"room_service"`
		Breakfast    int     `json:"breakfast"`
		Refund       int     `json:"refund"`
		GuestMax     int     `json:"guest_max"`
		RefundCharge float64 `json:"refund_charge"`
		RoomPitcure  string  `json:"room_picture"`
		RewardPoint  int64   `json:"reward_point"`
	}{
		HotelID:        a.HotelID,
		RoomType:       a.RoomType,
		RoomTypeName:   a.RoomTypeName,
		RoomOpen:       a.RoomOpen,
		AvailableRooms: a.AvailableRooms,
		Price:          a.Price,

		DiscPromo:      helper.NullInt64ToInt64(a.DiscPromo),
		DiscCampaign:   helper.NullInt64ToInt64(a.DiscCampaign),
		ValPromo:       a.ValPromo,
		ValCampaign:    a.ValCampaign,
		Taxes:          a.Taxes,
		DiscCharge:     a.DiscCharge,
		TotalMinCharge: a.TotalMinCharge,

		CheckinDate:  helper.TimeToString(a.CheckinDate),
		AC:           a.AC,
		Fan:          a.Fan,
		TV:           a.TV,
		Wifi:         a.Wifi,
		HotCool:      a.HotCool,
		Soap:         a.Soap,
		RoomService:  a.RoomService,
		Breakfast:    a.Breakfast,
		Refund:       a.Refund,
		GuestMax:     a.GuestMax,
		RefundCharge: a.RefundCharge,
		RoomPitcure:  a.RoomPitcure,
		RewardPoint:  a.RewardPoint,
	})
}

type SHotelCount struct {
	HotelID string `db:"hotel_id"`
	Count   int    `db:"count"`
}

func (a SHotelCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID string `json:"hotel_id"`
		Count   int    `json:"count"`
	}{
		HotelID: a.HotelID,
		Count:   a.Count,
	})
}

type SRoomCount struct {
	RoomName string `db:"room_type_name"`
	RoomType string `db:"room_type"`
	Count    int    `db:"count"`
}

func (a SRoomCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RoomType string `json:"room_type"`
		Count    int    `json:"count"`
	}{
		RoomType: a.RoomType,
		Count:    a.Count,
	})
}

type SHotel struct {
	HotelID        string          `db:"hotel_id"`
	HotelName      string          `db:"hotel_name"`
	RoadAddress    string          `db:"road_address"`
	MainPicture    string          `db:"main_picture"`
	Longitude      sql.NullString  `db:"long"`
	Latitude       sql.NullString  `db:"lat"`
	Stars          int             `db:"stars"`
	Description    string          `db:"description"`
	BookingType    int             `db:"booking_type"`
	HotelType      string          `db:"hotel_type"`
	RegencyID      string          `db:"regency_id"`
	ProvinceName   string          `db:"province_name"`
	MinCharge      int64           `db:"min_charge"`
	DiscPromo      sql.NullInt64   `db:"f_promo"`
	DiscCampaign   sql.NullInt64   `db:"f_campaign"`
	ValPromo       int64           `json:"promo_value"`
	ValCampaign    int64           `json:"campaign_value"`
	Taxes          int64           `json:"taxes"`
	DiscCharge     int64           `json:"min_charge_aftertax"`
	TotalMinCharge int64           `json:"total_min_charge"`
	Timezone       string          `db:"timezone"`
	AC             int             `db:"ac"`
	Fan            int             `db:"fan"`
	Tv             int             `db:"tv"`
	Wifi           int             `db:"wifi"`
	HotCool        int             `db:"hot_cool"`
	Soap           int             `db:"soap"`
	Reception      int             `db:"reception"`
	Resto          int             `db:"resto"`
	Pool           int             `db:"pool"`
	Gym            int             `db:"gym"`
	RoomService    int             `db:"room_service"`
	Parking        int             `db:"parking"`
	Breakfast      int             `db:"breakfast"`
	Refund         int             `db:"refund"`
	GuestMax       int             `db:"guest_max"`
	RefundCharge   sql.NullFloat64 `db:"refund_charge"`
	Review         float32         `db:"review"`
	AvailablePromo sql.NullFloat64 `db:"available_promo"`
	CheckInTime    string          `db:"check_in_time"`
	CheckOutTime   string          `db:"check_out_time"`
	AreaLoc        string          `db:"area_loc"`
	HotelFavorite  bool            `json:"hotel_favorite"`
}

func (a SHotel) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID        string  `json:"hotel_id"`
		HotelName      string  `json:"hotel_name"`
		RoadAddress    string  `json:"road_address"`
		MainPicture    string  `json:"main_picture"`
		Longitude      string  `json:"long"`
		Latitude       string  `json:"lat"`
		Stars          int     `json:"stars"`
		Description    string  `json:"description"`
		BookingType    int     `json:"booking_type"`
		HotelType      string  `json:"hotel_type"`
		RegencyID      string  `json:"regency_id"`
		RegencyName    string  `json:"regency_name"`
		ProvinceName   string  `json:"province_name"`
		MinCharge      int64   `json:"min_charge"`
		DiscPromo      int64   `json:"promo_disc"`
		DiscCampaign   int64   `json:"campaign_disc"`
		ValPromo       int64   `json:"promo_value"`
		ValCampaign    int64   `json:"campaign_value"`
		Taxes          int64   `json:"taxes"`
		DiscCharge     int64   `json:"min_charge_beforetax"`
		TotalMinCharge int64   `json:"total_min_charge"`
		Timezone       string  `json:"timezone"`
		AC             int     `json:"ac"`
		Fan            int     `json:"fan"`
		Tv             int     `json:"tv"`
		Wifi           int     `json:"wifi"`
		HotCool        int     `json:"hot_cool"`
		Soap           int     `json:"soap"`
		Reception      int     `json:"reception"`
		Resto          int     `json:"resto"`
		Pool           int     `json:"pool"`
		Gym            int     `json:"gym"`
		RoomService    int     `json:"room_service"`
		Parking        int     `json:"parking"`
		Breakfast      int     `json:"breakfast"`
		Refund         int     `json:"refund"`
		GuestMax       int     `json:"guest_max"`
		RefundCharge   float64 `json:"refund_charge"`
		Review         float32 `json:"review"`
		AvailablePromo float64 `json:"available_promo"`
		CheckInTime    string  `json:"check_in_time"`
		CheckOutTime   string  `json:"check_out_time"`
		AreaLoc        string  `json:"area_loc"`
		HotelFavorite  bool    `json:"hotel_favorite"`
	}{
		HotelID:        a.HotelID,
		HotelName:      a.HotelName,
		RoadAddress:    a.RoadAddress,
		MainPicture:    a.MainPicture,
		Longitude:      helper.NullStringToString(a.Longitude),
		Latitude:       helper.NullStringToString(a.Latitude),
		Stars:          a.Stars,
		Description:    a.Description,
		BookingType:    a.BookingType,
		HotelType:      a.HotelType,
		RegencyID:      a.RegencyID,
		ProvinceName:   a.ProvinceName,
		MinCharge:      a.MinCharge,
		DiscPromo:      helper.NullInt64ToInt64(a.DiscPromo),
		DiscCampaign:   helper.NullInt64ToInt64(a.DiscCampaign),
		ValPromo:       a.ValPromo,
		ValCampaign:    a.ValCampaign,
		Taxes:          a.Taxes,
		DiscCharge:     a.DiscCharge,
		TotalMinCharge: a.TotalMinCharge,
		Timezone:       a.Timezone,
		AC:             a.AC,
		Fan:            a.Fan,
		Tv:             a.Tv,
		Wifi:           a.Wifi,
		HotCool:        a.HotCool,
		Soap:           a.Soap,
		Reception:      a.Reception,
		Resto:          a.Resto,
		Pool:           a.Pool,
		Gym:            a.Gym,
		RoomService:    a.RoomService,
		Parking:        a.Parking,
		Breakfast:      a.Breakfast,
		Refund:         a.Refund,
		GuestMax:       a.GuestMax,
		RefundCharge:   helper.NullFloat64ToFloat64(a.RefundCharge),
		Review:         a.Review,
		AvailablePromo: helper.NullFloat64ToFloat64(a.AvailablePromo),
		CheckInTime:    a.CheckInTime,
		CheckOutTime:   a.CheckOutTime,
		AreaLoc:        a.AreaLoc,
		HotelFavorite:  a.HotelFavorite,
	})
}

type SHotels struct {
	NextPage bool     `json:"next_page"`
	Hotels   []SHotel `json:"hotels"`
}

func (a SHotels) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage bool     `json:"next_page"`
		Hotels   []SHotel `json:"hotels"`
	}{
		NextPage: a.NextPage,
		Hotels:   a.Hotels,
	})
}

type SHotel2 struct {
	HotelID         string          `db:"hotel_id"`
	HotelName       string          `db:"hotel_name"`
	RoadAddress     string          `db:"road_address"`
	MainPicture     string          `db:"main_picture"`
	Longitude       sql.NullString  `db:"long"`
	Latitude        sql.NullString  `db:"lat"`
	Stars           int             `db:"stars"`
	Description     string          `db:"description"`
	BookingType     int             `db:"booking_type"`
	HotelType       string          `db:"hotel_type"`
	RegencyID       string          `db:"regency_id"`
	ProvinceName    string          `db:"province_name"`
	MinCharge       int64           `db:"min_charge"`
	AreaLoc         string          `db:"area_loc"`
	Review          float32         `db:"review"`
	Timezone        string          `db:"timezone"`
	CheckInTime     string          `db:"check_in_time"`
	CheckOutTime    string          `db:"check_out_time"`
	Refund          int             `db:"refund"`
	GuestMax        int             `db:"guest_max"`
	RefundCharge    sql.NullFloat64 `db:"refund_charge"`
	DiscPromo       sql.NullInt64   `db:"promo"`
	DiscCampaign    sql.NullInt64   `db:"campaign"`
	ValPromo        int64           `json:"promo_value"`
	ValCampaign     int64           `json:"campaign_value"`
	Taxes           int64           `json:"taxes"`
	DiscCharge      int64           `json:"min_charge_aftertax"`
	TotalMinCharge  int64           `json:"total_min_charge"`
	AC              int             `db:"ac"`
	CarParking      int             `db:"car_parking"`
	BikeParking     int             `db:"bike_parking"`
	Wifi            int             `db:"wifi"`
	OutdoorPool     int             `db:"outdoor_pool"`
	IndoorPool      int             `db:"indoor_pool"`
	Gym             int             `db:"gym"`
	Receptionist    int             `db:"receptionist"`
	Limited_checkin int             `db:"limited_checkin"`
	Meeting_room    int             `db:"meeting_room"`
	Coffee_shop     int             `db:"coffee_shop"`
	Refrigerator    int             `db:"refrigerator"`
	A24h_checkin    int             `db:"24h_checkin"`
	Minibar         int             `db:"minibar"`
	Smoking_area    int             `db:"smoking_area"`
	Atm_centre      int             `db:"atm_centre"`
}

func (a SHotel2) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID         string  `json:"hotel_id"`
		HotelName       string  `json:"hotel_name"`
		RoadAddress     string  `json:"road_address"`
		MainPicture     string  `json:"main_picture"`
		Longitude       string  `json:"long"`
		Latitude        string  `json:"lat"`
		Stars           int     `json:"stars"`
		Description     string  `json:"description"`
		BookingType     int     `json:"booking_type"`
		HotelType       string  `json:"hotel_type"`
		RegencyID       string  `json:"regency_id"`
		RegencyName     string  `json:"regency_name"`
		ProvinceName    string  `json:"province_name"`
		MinCharge       int64   `json:"min_charge"`
		DiscPromo       int64   `json:"promo_disc"`
		DiscCampaign    int64   `json:"campaign_disc"`
		ValPromo        int64   `json:"promo_value"`
		ValCampaign     int64   `json:"campaign_value"`
		Taxes           int64   `json:"taxes"`
		DiscCharge      int64   `json:"min_charge_beforetax"`
		TotalMinCharge  int64   `json:"total_min_charge"`
		Timezone        string  `json:"timezone"`
		AC              int     `json:"ac"`
		Wifi            int     `json:"wifi"`
		Gym             int     `json:"gym"`
		Receptionist    int     `json:"receptionist"`
		Limited_checkin int     `json:"limited_checkin"`
		Meeting_room    int     `json:"meeting_room"`
		Coffee_shop     int     `json:"coffee_shop"`
		Refrigerator    int     `json:"refrigerator"`
		A24h_checkin    int     `json:"a24h_checkin"`
		Minibar         int     `json:"minibar"`
		Smoking_area    int     `json:"smoking_area"`
		Atm_centre      int     `json:"atm_centre"`
		Refund          int     `json:"refund"`
		GuestMax        int     `json:"guest_max"`
		RefundCharge    float64 `json:"refund_charge"`
		Review          float32 `json:"review"`
		CheckInTime     string  `json:"check_in_time"`
		CheckOutTime    string  `json:"check_out_time"`
		AreaLoc         string  `json:"area_loc"`
		CarParking      int     `json:"car_parking"`
		BikeParking     int     `json:"bike_parking"`
		OutdoorPool     int     `json:"outdoor_pool"`
		IndoorPool      int     `json:"indoor_pool"`
	}{
		HotelID:         a.HotelID,
		HotelName:       a.HotelName,
		RoadAddress:     a.RoadAddress,
		MainPicture:     a.MainPicture,
		Longitude:       helper.NullStringToString(a.Longitude),
		Latitude:        helper.NullStringToString(a.Latitude),
		Stars:           a.Stars,
		Description:     a.Description,
		BookingType:     a.BookingType,
		HotelType:       a.HotelType,
		RegencyID:       a.RegencyID,
		ProvinceName:    a.ProvinceName,
		MinCharge:       a.MinCharge,
		DiscPromo:       helper.NullInt64ToInt64(a.DiscPromo),
		DiscCampaign:    helper.NullInt64ToInt64(a.DiscCampaign),
		ValPromo:        a.ValPromo,
		ValCampaign:     a.ValCampaign,
		Taxes:           a.Taxes,
		DiscCharge:      a.DiscCharge,
		TotalMinCharge:  a.TotalMinCharge,
		Timezone:        a.Timezone,
		AC:              a.AC,
		Wifi:            a.Wifi,
		Gym:             a.Gym,
		Refund:          a.Refund,
		GuestMax:        a.GuestMax,
		RefundCharge:    helper.NullFloat64ToFloat64(a.RefundCharge),
		Review:          a.Review,
		CheckInTime:     a.CheckInTime,
		CheckOutTime:    a.CheckOutTime,
		AreaLoc:         a.AreaLoc,
		CarParking:      a.CarParking,
		BikeParking:     a.BikeParking,
		OutdoorPool:     a.OutdoorPool,
		IndoorPool:      a.IndoorPool,
		Receptionist:    a.Receptionist,
		Limited_checkin: a.Limited_checkin,
		Meeting_room:    a.Meeting_room,
		Coffee_shop:     a.Coffee_shop,
		Refrigerator:    a.Refrigerator,
		A24h_checkin:    a.A24h_checkin,
		Minibar:         a.Minibar,
		Smoking_area:    a.Smoking_area,
		Atm_centre:      a.Atm_centre,
	})
}

type SHotels2 struct {
	NextPage bool      `json:"next_page"`
	Hotels   []SHotel2 `json:"hotels"`
}

func (a SHotels2) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage bool      `json:"next_page"`
		Hotels   []SHotel2 `json:"hotels"`
	}{
		NextPage: a.NextPage,
		Hotels:   a.Hotels,
	})
}

type SHotelsOptionalFilter struct {
	MaxPrice  float64  `json:"max_price"`
	MinPrice  float64  `json:"min_price"`
	Stars     []string `json:"stars"`
	Facility  []string `json:"facility"`
	HotelType []string `json:"hotel_type"`
}

func (a SHotelsOptionalFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		MaxPrice  float64 `json:"max_price"`
		MinPrice  float64 `json:"min_price"`
		Stars     string  `json:"stars"`
		Facility  string  `json:"facility"`
		HotelType string  `json:"hotel_type"`
	}{
		MaxPrice:  a.MaxPrice,
		MinPrice:  a.MinPrice,
		Stars:     helper.StringToArrayString(a.Stars),
		Facility:  helper.StringToArrayString(a.Facility),
		HotelType: helper.StringToArrayString(a.HotelType),
	})
}

type SFacilityM struct {
	HotelID     int `db:"hotel_id"`
	AC          int `db:"ac"`
	Fan         int `db:"fan"`
	TV          int `db:"tv"`
	Wifi        int `db:"wifi"`
	HotCool     int `db:"hot_cool"`
	Soap        int `db:"soap"`
	Reception   int `db:"reception"`
	Resto       int `db:"resto"`
	Pool        int `db:"pool"`
	Gym         int `db:"gym"`
	RoomService int `db:"room_service"`
	Parking     int `db:"parking"`
	Breakfast   int `db:"breakfast"`
	Refund      int `db:"refund"`
	GuestMax    int `db:"guest_max"`
}

func (a SFacilityM) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelID     int `json:"hotel_id"`
		AC          int `json:"ac"`
		Fan         int `json:"fan"`
		TV          int `json:"tv"`
		Wifi        int `json:"wifi"`
		HotCool     int `json:"hot_cool"`
		Soap        int `json:"soap"`
		Reception   int `json:"reception"`
		Resto       int `json:"resto"`
		Pool        int `json:"pool"`
		Gym         int `json:"gym"`
		RoomService int `json:"room_service"`
		Parking     int `json:"parking"`
		Breakfast   int `json:"breakfast"`
		Refund      int `json:"refund"`
		GuestMax    int `json:"guest_max"`
	}{
		HotelID:     a.HotelID,
		AC:          a.AC,
		Fan:         a.Fan,
		TV:          a.TV,
		Wifi:        a.Wifi,
		HotCool:     a.HotCool,
		Soap:        a.Soap,
		Reception:   a.Reception,
		Resto:       a.Resto,
		Pool:        a.Pool,
		Gym:         a.Gym,
		RoomService: a.RoomService,
		Parking:     a.Parking,
		Breakfast:   a.Breakfast,
		Refund:      a.Refund,
		GuestMax:    a.GuestMax,
	})
}

type SHotelTypeM struct {
	HotelTypeID   string `db:"id"`
	HotelTypeDesc string `db:"type_desc"`
}

func (a SHotelTypeM) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HotelTypeID   string `json:"hotel_type_id"`
		HotelTypeDesc string `json:"hotel_type_desc"`
	}{
		HotelTypeID:   a.HotelTypeID,
		HotelTypeDesc: a.HotelTypeDesc,
	})
}

type SDHotelPromosApi struct {
	NextPage bool            `json:"next_page"`
	Promo    []SDHotelPromos `json:"promo"`
}

func (a SDHotelPromosApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage bool            `json:"next_page"`
		Promo    []SDHotelPromos `json:"promo"`
	}{
		NextPage: a.NextPage,
		Promo:    a.Promo,
	})
}

type SDHotelPromos struct {
	PromoID          string          `db:"promo_id"`
	PromoName        string          `db:"promo_name"`
	TS               time.Time       `db:"ts"`
	StartDate        time.Time       `db:"start_date"`
	EndDate          time.Time       `db:"end_date"`
	Term             string          `db:"terms"`
	Description      string          `db:"description"`
	FilePart1        sql.NullString  `db:"file_part_1"`
	FilePart2        sql.NullString  `db:"file_part_2"`
	DateOfMonth      string          `db:"date_of_month"`
	DayOfWeek        string          `db:"day_of_week"`
	LinkedCampaign   string          `db:"linked_campaign"`
	Owner            string          `db:"owner"`
	HotelID          string          `db:"hotel_id"`
	RoomType         int             `db:"room_type"`
	PromoVal         float64         `db:"promo_val"`
	CampaignPromoVal sql.NullFloat64 `db:"campaign_promo_val"`
	TimeOfDayStart   sql.NullString  `db:"time_of_day_start"`
	TimeOfDayEnd     sql.NullString  `db:"time_of_day_end"`
}

func (a SDHotelPromos) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PromoID          string  `json:"promo_id"`
		PromoName        string  `json:"promo_name"`
		TS               string  `json:"ts"`
		StartDate        string  `json:"start_date"`
		EndDate          string  `json:"end_date"`
		Term             string  `json:"term"`
		Description      string  `json:"description"`
		FilePart1        string  `json:"file_part_1"`
		FilePart2        string  `json:"file_part_2"`
		DateOfMonth      string  `json:"date_of_month"`
		DayOfWeek        string  `json:"day_of_week"`
		LinkedCampaign   string  `json:"linked_campaign"`
		Owner            string  `json:"owner"`
		HotelID          string  `json:"hotel_id"`
		RoomType         int     `json:"room_type"`
		PromoVal         float64 `json:"promo_val"`
		CampaignPromoVal float64 `json:"campaign_promo_val"`
		TimeOfDayStart   string  `json:"time_of_day_start"`
		TimeOfDayEnd     string  `json:"time_of_day_end"`
	}{
		PromoID:          a.PromoID,
		PromoName:        a.PromoName,
		TS:               helper.TimeToString(a.TS),
		StartDate:        helper.TimeToString(a.StartDate),
		EndDate:          helper.TimeToString(a.EndDate),
		Term:             a.Term,
		Description:      a.Description,
		FilePart1:        helper.NullStringToString(a.FilePart1),
		FilePart2:        helper.NullStringToString(a.FilePart2),
		DateOfMonth:      a.DateOfMonth,
		DayOfWeek:        a.DayOfWeek,
		LinkedCampaign:   a.LinkedCampaign,
		Owner:            a.Owner,
		HotelID:          a.HotelID,
		RoomType:         a.RoomType,
		PromoVal:         a.PromoVal,
		CampaignPromoVal: helper.NullFloat64ToFloat64(a.CampaignPromoVal),
		TimeOfDayStart:   helper.NullStringToString(a.TimeOfDayStart),
		TimeOfDayEnd:     helper.NullStringToString(a.TimeOfDayEnd),
	})
}

type SHotelAutocomplete struct {
	LookUpKey        string `db:"lookup_key"`
	AutoCompleteType string `db:"code"`
}

func (a SHotelAutocomplete) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		LookUpKey        string `json:"lookup_key"`
		AutoCompleteType string `json:"result_type"`
	}{
		LookUpKey:        a.LookUpKey,
		AutoCompleteType: a.AutoCompleteType,
	})
}
