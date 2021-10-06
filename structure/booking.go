package structure

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/ruangnyaman/rna-ecommerce-backend/helper"
)

type SBookingGuestName struct {
	Name []string `json:"name"`
}
type SBookings struct {
	NextPage       bool       `json:"next_page"`
	BookingHistory []SBooking `json:"booking_history"`
}

func (a SBookings) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage       bool       `json:"next_page"`
		BookingHistory []SBooking `json:"booking_history"`
	}{
		NextPage:       a.NextPage,
		BookingHistory: a.BookingHistory,
	})
}

type SBooking struct {
	BookingID     string    `db:"booking_id"`
	UserID        string    `db:"user_id"`
	BookingType   int       `db:"booking_type"`
	Description   string    `db:"description"`
	CreatedAt     time.Time `db:"created_at"`
	TotalCharges  string    `db:"total_charge"`
	BookingStatus int       `db:"booking_status"`
	PaymentStatus int       `db:"payment_status"`
	PaymentDue    time.Time `db:"payment_due"`
	BankName      string    `db:"to_bank_name"`
	PaymentMethod int       `db:"payment_method"`
}

func (a SBooking) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		BookingID     string `json:"booking_id"`
		UserID        string `json:"user_id"`
		BookingType   int    `json:"booking_type"`
		Description   string `json:"description"`
		CreatedAt     string `json:"created_at"`
		TotalCharges  string `json:"total_charge"`
		BookingStatus int    `json:"booking_status"`
		PaymentStatus int    `json:"payment_status"`
		PaymentDue    string `json:"payment_due"`
		BankName      string `json:"to_bank_name"`
		PaymentMethod int    `json:"payment_method"`
	}{
		BookingID:     a.BookingID,
		UserID:        a.UserID,
		BookingType:   a.BookingType,
		Description:   a.Description,
		CreatedAt:     helper.TimeToString(a.CreatedAt),
		TotalCharges:  a.TotalCharges,
		BookingStatus: a.BookingStatus,
		PaymentStatus: a.PaymentStatus,
		PaymentDue:    helper.TimeToString(a.PaymentDue),
		BankName:      a.BankName,
		PaymentMethod: a.PaymentMethod,
	})
}

type SBookingActivity struct {
	ActivityBookingID string        `db:"activity_booking_id"`
	UserID            string        `db:"user_id"`
	TS                time.Time     `db:"ts"`
	ActivityID        string        `db:"activity_id"`
	ActivityName      string        `db:"activity_name"`
	ActivityPackageID string        `db:"activity_package_id"`
	TicketQty         int           `db:"ticket_qty"`
	DiscountValue     int64         `db:"discount_val"`
	Price             int64         `db:"price"`
	Tax               int64         `db:"tax"`
	TotalCharge       string        `db:"total_charge"`
	Status            int           `db:"status"`
	PaymentStatus     int           `db:"payment_status"`
	ActivityType      int           `db:"activity_type"`
	PointRewards      sql.NullInt64 `db:"point_rewards"`
	TicketDate        sql.NullTime  `db:"ticket_date"`
	PaymentAcc        string        `db:"va_no"`
}

func (a SBookingActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ActivityBookingID string `json:"activity_booking_id"`
		UserID            string `json:"user_id"`
		TS                string `json:"ts"`
		ActivityID        string `json:"activity_id"`
		ActivityName      string `json:"activity_name"`
		ActivityPackageID string `json:"activity_package_id"`
		TicketQty         int    `json:"ticket_qty"`
		DiscountValue     int64  `json:"discount_val"`
		Price             int64  `json:"price"`
		Tax               int64  `json:"tax"`
		TotalCharge       string `json:"total_charge"`
		Status            int    `json:"status"`
		PaymentStatus     int    `json:"payment_status"`
		ActivityType      int    `json:"activity_type"`
		PointRewards      int64  `json:"point_rewards"`
		TicketDate        string `json:"ticket_date"`
		PaymentAcc        string `json:"acc_number"`
	}{
		ActivityBookingID: a.ActivityBookingID,
		UserID:            a.UserID,
		TS:                helper.TimeToString(a.TS),
		ActivityID:        a.ActivityID,
		ActivityName:      a.ActivityName,
		ActivityPackageID: a.ActivityPackageID,
		TicketQty:         a.TicketQty,
		DiscountValue:     a.DiscountValue,
		Price:             a.Price,
		Tax:               a.Tax,
		TotalCharge:       a.TotalCharge,
		Status:            a.Status,
		PaymentStatus:     a.PaymentStatus,
		ActivityType:      a.ActivityType,
		PointRewards:      helper.NullInt64ToInt64(a.PointRewards),
		TicketDate:        helper.NullTimeToString(a.TicketDate),
		PaymentAcc:        a.PaymentAcc,
	})
}

type SBookingRoom struct {
	BookingID      string         `db:"booking_id"`
	CreatedAt      time.Time      `db:"created_at"`
	UserID         string         `db:"user_id"`
	HotelName      string         `db:"hotel_name"`
	RoomTypeName   string         `db:"room_type_name"`
	CheckinDate    time.Time      `db:"checkin_date"`
	CheckOutDate   time.Time      `db:"checkout_date"`
	TotalPrice     int64          `db:"total_price"`
	PromoVal       int64          `db:"promo_val"`
	TotalCharges   string         `db:"total_charge"`
	BookingStatus  int32          `db:"booking_status"`
	PaymentStatus  int32          `db:"payment_status"`
	GuestName      string         `db:"quest_name"`
	SpecialRequest sql.NullString `db:"special_request"`
	HotelID        string         `db:"hotel_id"`
	RoomTypeID     string         `db:"room_type_id"`
	RoomQty        int            `db:"room_qty"`
	TaxAmount      int            `db:"tax_amnt"`
	PointRewards   sql.NullInt64  `db:"point_rewards"`
	PaymentAcc     string         `db:"acc_number"`
}

func (a SBookingRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		BookingID      string `json:"booking_id"`
		CreatedAt      string `json:"created_at"`
		UserID         string `json:"user_id"`
		HotelName      string `json:"hotel_name"`
		RoomTypeName   string `json:"room_type_name"`
		CheckinDate    string `json:"checkin_date"`
		CheckOutDate   string `json:"checkout_date"`
		TotalPrice     int64  `json:"total_price"`
		PromoVal       int64  `json:"promo_val"`
		TotalCharges   string `json:"total_charge"`
		BookingStatus  int32  `json:"booking_status"`
		PaymentStatus  int32  `json:"payment_status"`
		GuestName      string `json:"quest_name"`
		SpecialRequest string `json:"special_request"`
		HotelID        string `json:"hotel_id"`
		RoomTypeID     string `json:"room_type_id"`
		RoomQty        int    `json:"room_qty"`
		TaxAmount      int    `json:"tax_amnt"`
		PointRewards   int64  `json:"point_rewards"`
		PaymentAcc     string `json:"acc_number"`
	}{
		BookingID:      a.BookingID,
		CreatedAt:      helper.TimeToString(a.CreatedAt),
		UserID:         a.UserID,
		HotelName:      a.HotelName,
		RoomTypeName:   a.RoomTypeName,
		CheckinDate:    helper.TimeToString(a.CheckinDate),
		CheckOutDate:   helper.TimeToString(a.CheckOutDate),
		TotalPrice:     a.TotalPrice,
		PromoVal:       a.PromoVal,
		TotalCharges:   a.TotalCharges,
		BookingStatus:  a.BookingStatus,
		PaymentStatus:  a.PaymentStatus,
		GuestName:      a.GuestName,
		SpecialRequest: helper.NullStringToString(a.SpecialRequest),
		HotelID:        a.HotelID,
		RoomTypeID:     a.RoomTypeID,
		RoomQty:        a.RoomQty,
		TaxAmount:      a.TaxAmount,
		PointRewards:   helper.NullInt64ToInt64(a.PointRewards),
		PaymentAcc:     a.PaymentAcc,
	})
}

type SPreOrderDetail struct {
	RoomTypeName string    `db:"room_type_name"`
	TotalPrice   float64   `db:"total_price"`
	Date         time.Time `db:"checkin_date"`
	RewardPoint  int64     `db:"reward_point"`
}

type SPreOrderCalc struct {
	RoomTypeName   string  `db:"room_type_name"`
	TotalPrice     float64 `json:"total_price"`
	DiscPromo      int64   `json:"promo_disc"`
	DiscCampaign   int64   `json:"campaign_disc"`
	ValPromo       float64 `json:"promo_value"`
	ValCampaign    float64 `json:"campaign_value"`
	SubTotalCharge float64 `json:"sub_total"`
	TaxesAndFee    float64 `json:"taxes_and_fee"`
	TotalCharge    float64 `json:"total_charge"`
	HotelInfo      SHotel  `db:"hotel_info"`
	OrderDetail    []SPreOrderDetail
	RewardPoint    int64 `db:"reward_point"`
}
type SPreOrder struct {
	RoomTypeName string `json:"room_type_name"`
	Stay         int    `json:"stay"`
	Room         string `json:"room_total"`

	DiscPromo      int64   `json:"promo_disc"`
	DiscCampaign   int64   `json:"campaign_disc"`
	ValPromo       float64 `json:"promo_value"`
	ValCampaign    float64 `json:"campaign_value"`
	SubTotalCharge float64 `json:"sub_total"`
	TotalPrice     float64 `json:"total_price"`
	TaxesAndFee    float64 `json:"taxes_and_fee"`
	TotalCharge    float64 `json:"total_charge"`

	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	HotelInfo   SHotel            `json:"hotel_info"`
	OrderDetail []SPreOrderDetail `json:"rooms_detail"`
	RewardPoint int64             `json:"reward_point"`
}

func (a SPreOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RoomTypeName string `json:"room_type_name"`
		Stay         int    `json:"stay"`
		Room         string `json:"room_total"`

		DiscPromo      int64   `json:"promo_disc"`
		DiscCampaign   int64   `json:"campaign_disc"`
		ValPromo       float64 `json:"promo_value"`
		ValCampaign    float64 `json:"campaign_value"`
		SubTotalCharge float64 `json:"sub_total"`
		TotalPrice     float64 `json:"total_price"`
		TaxesAndFee    float64 `json:"taxes_and_fee"`
		TotalCharge    float64 `json:"total_charge"`

		Name        string            `json:"name"`
		Email       string            `json:"email"`
		Phone       string            `json:"phone"`
		HotelInfo   SHotel            `json:"hotel_info"`
		OrderDetail []SPreOrderDetail `json:"rooms_detail"`
		RewardPoint int64             `json:"reward_point"`
	}{
		RoomTypeName:   a.RoomTypeName,
		Stay:           a.Stay,
		Room:           a.Room,
		DiscPromo:      a.DiscPromo,
		DiscCampaign:   a.DiscCampaign,
		ValPromo:       a.ValPromo,
		ValCampaign:    a.ValCampaign,
		SubTotalCharge: a.SubTotalCharge,
		TotalPrice:     a.TotalPrice,
		TaxesAndFee:    a.TaxesAndFee,
		TotalCharge:    a.TotalCharge,
		Name:           a.Name,
		Email:          a.Email,
		Phone:          a.Phone,
		HotelInfo:      a.HotelInfo,
		OrderDetail:    a.OrderDetail,
		RewardPoint:    a.RewardPoint,
	})
}

type SBookingPrice struct {
	Price        float64 `db:"total_price"`
	RewardPoint  int64   `db:"total_reward_point"`
	HotelAddress string  `db:"road_address"`
}
type SParamHolderForCheckBooking struct {
	HotelID        string
	PaymentID      string
	RoomTypeID     string
	DurationString string
	RoomQty        float64
}
type SOutputHolderForCheckBooking struct {
	HotelName     string
	BankName      string
	PaymentType   string
	TaxAmount     float64
	TotalPrice    float64
	TotalCharge   int64
	DiscountValue float64
	RewardPoint   int64
	PromoID       sql.NullString
	CheckinTime   string
	CheckOutTime  string
}
type SParamHolderForBooking struct {
	HotelID        string
	RoomTypeID     string
	RoomQty        float64
	UserID         string
	CheckinDate    time.Time
	CheckOutDate   time.Time
	GuestName      string
	SpecialRequest sql.NullString
	GuestNum       int

	// DiscountValue  float64
	// TotalPrice     float64
	// TaxAmount      float64
	// TotalCharge    int64
	// PromoID        sql.NullString
	// RewardPoint    int64
	// BankName       string

	FullName       string
	DueDate        time.Time
	Payment        SPaymentMethod
	DurationString string
	DurationInt    int
}

type SOutputHolderForBooking struct {
	BookingID    string
	RoomTypeName string
	AccNumber    string
	TotalCharge  int64
	HotelAddress string
}

type SPreOrderEvent struct {
	PackageDetail SEventPackage `json:"event_package"`
	TotalPrice    int64         `json:"total_price"`
	Discount      int64         `json:"discount"`
	SubTotal      int64         `json:"sub_total"`
	TaxesAndFee   float64       `json:"taxes_and_fee"`
	TotalCharge   int64         `json:"total_charge"`
	PointRewards  int64         `json:"point_rewards"`
}

func (a SPreOrderEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PackageDetail SEventPackage `json:"event_package"`
		TotalPrice    int64         `json:"total_price"`
		Discount      int64         `json:"discount"`
		SubTotal      int64         `json:"sub_total"`
		TaxesAndFee   float64       `json:"taxes_and_fee"`
		TotalCharge   int64         `json:"total_charge"`
		PointRewards  int64         `json:"point_rewards"`
	}{
		PackageDetail: a.PackageDetail,
		TotalPrice:    a.TotalPrice,
		Discount:      a.Discount,
		SubTotal:      a.SubTotal,
		TaxesAndFee:   a.TaxesAndFee,
		TotalCharge:   a.TotalCharge,
		PointRewards:  a.PointRewards,
	})
}

type SParamHolderForBookingActivity struct {
	ActivityPackageID string
	UserID            string
	Qty               int64
	Taxes             int
	Payment           SPaymentMethod
	DueDate           time.Time
	Name              string
	FullName          string
	TicketDate        sql.NullTime
}

type SOutputHolderForBookingActivity struct {
	ActivityPackageID   string
	ActivityID          string
	ActicityBookingID   string
	ActivityPackageName string
	TotalCharges        int64
	BankName            string
	PaymentAcc          string
	PaymentType         string
}

// Activiry
type SPreOrderTodo struct {
	PackageDetail STodoPackage `json:"event_package"`
	TotalPrice    int64        `json:"total_price"`
	Discount      int64        `json:"discount"`
	SubTotal      int64        `json:"sub_total"`
	TaxesAndFee   float64      `json:"taxes_and_fee"`
	TotalCharge   int64        `json:"total_charge"`
	PointRewards  int64        `json:"point_rewards"`
}

func (a SPreOrderTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PackageDetail STodoPackage `json:"event_package"`
		TotalPrice    int64        `json:"total_price"`
		Discount      int64        `json:"discount"`
		SubTotal      int64        `json:"sub_total"`
		TaxesAndFee   float64      `json:"taxes_and_fee"`
		TotalCharge   int64        `json:"total_charge"`
		PointRewards  int64        `json:"point_rewards"`
	}{
		PackageDetail: a.PackageDetail,
		TotalPrice:    a.TotalPrice,
		Discount:      a.Discount,
		SubTotal:      a.SubTotal,
		TaxesAndFee:   a.TaxesAndFee,
		TotalCharge:   a.TotalCharge,
		PointRewards:  a.PointRewards,
	})
}

type SSpecialRequest struct {
	NonSmoking bool   `json:"non_smoking"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
	Other      string `json:"other"`
}
