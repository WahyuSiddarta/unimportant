package structure

import (
	"encoding/json"
	"time"

	"github.com/WahyuSiddarta/unimportant/helper"
)

type SActivitiesApi struct {
	NextPage bool          `json:"next_page"`
	Activity []SActivities `json:"activity"`
}

func (a SActivitiesApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage bool          `json:"next_page"`
		Activity []SActivities `json:"event"`
	}{
		NextPage: a.NextPage,
		Activity: a.Activity,
	})
}

type SActivities struct {
	TodoID           string    `db:"todo_id"`
	PartnerID        string    `db:"partner_id"`
	PatnerName       string    `db:"partner_name"`
	TS               time.Time `db:"ts"`
	LastUpdate       time.Time `db:"last_update"`
	TodoName         string    `db:"todo_name"`
	PlaceName        string    `db:"place_name"`
	AddressLoc       string    `db:"address_loc"`
	OpenClose        string    `db:"open_close"`
	AreaLoc          string    `db:"area_loc"`
	TodoCategoryID   string    `db:"todo_category_id"`
	TodoCategoryName string    `db:"todo_category_name"`
	Description      string    `db:"description"`
	Term             string    `db:"term"`
	Direction        string    `db:"direction"`
	FilePart         string    `db:"file_part"`
	ReviewScore      int64     `db:"review_score"`
	TicketDue        time.Time `db:"ticket_due"`
	Timezone         string    `db:"timezone"`
	PackageIncluding string    `db:"package_including"`
	Highlight        string    `db:"higlight"`
	TicketExchange   string    `db:"ticket_exchange"`
	Price            int64     `db:"price"`
}

func (a SActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TodoID           string `json:"todo_id"`
		PartnerID        string `json:"partner_id"`
		PatnerName       string `json:"partner_name"`
		TS               string `json:"ts"`
		LastUpdate       string `json:"last_update"`
		TodoName         string `json:"todo_name"`
		PlaceName        string `json:"place_name"`
		AddressLoc       string `json:"address_loc"`
		OpenClose        string `json:"open_close"`
		AreaLoc          string `json:"area_loc"`
		TodoCategoryID   string `json:"todo_category_id"`
		TodoCategoryName string `json:"todo_category_name"`
		Description      string `json:"description"`
		Term             string `json:"term"`
		Direction        string `json:"direction"`
		FilePart         string `json:"file_part"`
		ReviewScore      int64  `json:"review_score"`
		TicketDue        string `json:"ticket_due"`
		Timezone         string `json:"timezone"`
		PackageIncluding string `json:"package_including"`
		Highlight        string `json:"highlight"`
		TicketExchange   string `json:"ticket_exchange"`
		Price            int64  `json:"price"`
	}{
		TodoID:           a.TodoID,
		PartnerID:        a.PartnerID,
		PatnerName:       a.PatnerName,
		TS:               helper.TimeToString(a.TS),
		LastUpdate:       helper.TimeToString(a.LastUpdate),
		TodoName:         a.TodoName,
		PlaceName:        a.PlaceName,
		AddressLoc:       a.AddressLoc,
		OpenClose:        a.OpenClose,
		AreaLoc:          a.AreaLoc,
		TodoCategoryID:   a.TodoCategoryID,
		TodoCategoryName: a.TodoCategoryName,
		Description:      a.Description,
		Term:             a.Term,
		Direction:        a.Direction,
		FilePart:         a.FilePart,
		ReviewScore:      a.ReviewScore,
		TicketDue:        helper.TimeToString(a.TicketDue),
		Timezone:         a.Timezone,
		PackageIncluding: a.PackageIncluding,
		Highlight:        a.Highlight,
		TicketExchange:   a.TicketExchange,
		Price:            a.Price,
	})
}

type STodoPackage struct {
	TodoPackageID      string    `db:"todo_package_id"`
	TodoID             string    `db:"todo_id"`
	TS                 time.Time `db:"ts"`
	LastUpdate         time.Time `db:"last_update"`
	PackageName        string    `db:"package_name"`
	PackageDescription string    `db:"package_description"`
	PackageInformation string    `db:"package_information"`
	PackageTerm        string    `db:"package_term"`
	Refund             int64     `db:"refund"`
	Remaining          int64     `db:"remaining"`
	RealPrice          int64     `db:"real_price"`
	SellingPrice       int64     `db:"selling_price"`
	PointRewards       int64     `db:"point_rewards"`
	TicketDue          time.Time `db:"ticket_due"`
}

func (a STodoPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TodoPackageID      string `json:"todo_package_id"`
		TodoID             string `json:"todo_id"`
		TS                 string `json:"ts"`
		LastUpdate         string `json:"last_update"`
		PackageName        string `json:"package_name"`
		PackageDescription string `json:"package_description"`
		PackageInformation string `json:"package_information"`
		PackageTerm        string `json:"package_term"`
		Refund             int64  `json:"refund"`
		Remaining          int64  `json:"remaining"`
		RealPrice          int64  `json:"real_price"`
		SellingPrice       int64  `json:"selling_price"`
		PointRewards       int64  `json:"point_rewards"`
		TicketDue          string `json:"ticket_due"`
	}{
		TodoPackageID:      a.TodoPackageID,
		TodoID:             a.TodoID,
		TS:                 helper.TimeToString(a.TS),
		LastUpdate:         helper.TimeToString(a.LastUpdate),
		PackageName:        a.PackageName,
		PackageDescription: a.PackageDescription,
		PackageInformation: a.PackageInformation,
		PackageTerm:        a.PackageTerm,
		Refund:             a.Refund,
		Remaining:          a.Remaining,
		RealPrice:          a.RealPrice,
		SellingPrice:       a.SellingPrice,
		PointRewards:       a.PointRewards,
		TicketDue:          helper.TimeToString(a.TicketDue),
	})
}

type SEventListApi struct {
	NextPage bool         `json:"next_page"`
	Event    []SEventList `json:"event"`
}

func (a SEventListApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NextPage bool         `json:"next_page"`
		Event    []SEventList `json:"event"`
	}{
		NextPage: a.NextPage,
		Event:    a.Event,
	})
}

type SEventList struct {
	EventID         string    `db:"event_id"`
	PatnerID        string    `db:"partner_id"`
	PatnerName      string    `db:"partner_name"`
	TS              time.Time `db:"ts"`
	LastUpdate      time.Time `db:"last_update"`
	EventName       string    `db:"event_name"`
	PlaceLoc        string    `db:"place_loc"`
	AreaLoc         string    `db:"area_loc"`
	EventCategoryID string    `db:"event_category_id"`
	CategoryName    string    `db:"category_name"`
	PeriodStart     time.Time `db:"period_start"`
	PeriodEnd       time.Time `db:"period_end"`
	Description     string    `db:"description"`
	Term            string    `db:"term"`
	AddDesc         string    `db:"add_desc"`
	FilePart        string    `db:"file_part"`
	EventInfo       string    `db:"event_info"`
	EventDate       time.Time `db:"event_date"`
	TimeZone        string    `db:"timezone"`
	EventPrice      int64     `db:"event_price"`
}

func (a SEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		EventID         string `json:"event_id"`
		PatnerID        string `json:"partner_id"`
		PatnerName      string `json:"partner_name"`
		TS              string `json:"ts"`
		LastUpdate      string `json:"last_update"`
		EventName       string `json:"event_name"`
		PlaceLoc        string `json:"place_loc"`
		AreaLoc         string `json:"area_loc"`
		EventCategoryID string `json:"event_category_id"`
		CategoryName    string `json:"category_name"`
		PeriodStart     string `json:"period_start"`
		PeriodEnd       string `json:"period_end"`
		Description     string `json:"description"`
		Term            string `json:"term"`
		AddDesc         string `json:"add_desc"`
		FilePart        string `json:"file_part"`
		EventInfo       string `json:"event_info"`
		EventDate       string `json:"event_date"`
		TimeZone        string `json:"timezone"`
		EventPrice      int64  `json:"event_price"`
	}{
		EventID:         a.EventID,
		PatnerID:        a.PatnerID,
		PatnerName:      a.PatnerName,
		TS:              helper.TimeToString(a.TS),
		LastUpdate:      helper.TimeToString(a.LastUpdate),
		EventName:       a.EventName,
		PlaceLoc:        a.PlaceLoc,
		AreaLoc:         a.AreaLoc,
		EventCategoryID: a.EventCategoryID,
		CategoryName:    a.CategoryName,
		PeriodStart:     helper.TimeToString(a.PeriodStart),
		PeriodEnd:       helper.TimeToString(a.PeriodEnd),
		Description:     a.Description,
		Term:            a.Term,
		AddDesc:         a.AddDesc,
		FilePart:        a.FilePart,
		EventInfo:       a.EventInfo,
		EventDate:       helper.TimeToString(a.EventDate),
		TimeZone:        a.TimeZone,
		EventPrice:      a.EventPrice,
	})
}

type SEventPackage struct {
	PackageID          string    `db:"package_id"`
	EventID            string    `db:"event_id"`
	TS                 time.Time `db:"ts"`
	LastUpdate         time.Time `db:"last_update"`
	PackageName        string    `db:"package_name"`
	PackageDescription string    `db:"package_description"`
	PackageInformation string    `db:"package_information"`
	PackageTerm        string    `db:"package_term"`
	PackagePrice       int64     `db:"package_price"`
	Refund             int64     `db:"refund"`
	Remaining          int64     `db:"remaining"`
	RealPrice          int64     `db:"real_price"`
	PointRewards       int64     `db:"point_rewards"`
}

func (a SEventPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PackageID          string `json:"package_id"`
		EventID            string `json:"event_id"`
		TS                 string `json:"ts"`
		LastUpdate         string `json:"last_update"`
		PackageName        string `json:"package_name"`
		PackageDescription string `json:"package_description"`
		PackageInformation string `json:"package_information"`
		PackageTerm        string `json:"package_term"`
		PackagePrice       int64  `json:"package_price"`
		Refund             int64  `json:"refund"`
		Remaining          int64  `json:"remaining"`
		RealPrice          int64  `json:"real_price"`
		PointRewards       int64  `json:"point_rewards"`
	}{
		PackageID:          a.PackageID,
		EventID:            a.EventID,
		TS:                 helper.TimeToString(a.TS),
		LastUpdate:         helper.TimeToString(a.LastUpdate),
		PackageName:        a.PackageName,
		PackageDescription: a.PackageDescription,
		PackageInformation: a.PackageInformation,
		PackageTerm:        a.PackageTerm,
		PackagePrice:       a.PackagePrice,
		Refund:             a.Refund,
		Remaining:          a.Remaining,
		RealPrice:          a.RealPrice,
		PointRewards:       a.PointRewards,
	})
}
