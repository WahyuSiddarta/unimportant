package bni_helper

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ruangnyaman/rna-ecommerce-backend/conf"
	bniEnc "github.com/ruangnyaman/rna-ecommerce-backend/encrypt/bni"
	"github.com/ruangnyaman/rna-ecommerce-backend/helper"
	myHTTP "github.com/ruangnyaman/rna-ecommerce-backend/helper/http"
)

const (
	createBillingStr   = "createbilling"
	inquiryBillingStr  = "inquirybilling"
	updateBillingStr   = "updatebilling"
	billingOpenCode    = "o"
	billingOpenMinCode = "n"
	billingPartialCode = "i"
	billingFixCode     = "c"
	billingOpenMaxCode = "x"
)

var (
	bniSecret   string
	bniClientID string
	bniPrefixVA string
	trxURL      string
)

func init() {
	Config, err := conf.GetConfig()
	if err != nil {
		panic("Error opening config file")
	}

	bniSecret = Config.GetString("payment.provider_bni.secret")
	bniClientID = Config.GetString("payment.provider_bni.client_id")
	bniPrefixVA = Config.GetString("payment.provider_bni.prefix_va")
	trxURL = Config.GetString("payment.provider_bni.trx_url")
}

type SBNICreateBilling struct {
	Type            string `json:"type,omitempty"`          // mandatory
	ClientID        string `json:"client_id,omitempty"`     // mandatory
	TrxID           string `json:"trx_id,omitempty"`        // mandatory
	TrxAmount       string `json:"trx_amount,omitempty"`    // mandatory
	BillingType     string `json:"billing_type,omitempty"`  // mandatory
	CustomerName    string `json:"customer_name,omitempty"` // mandatory
	CustomerEmail   string `json:"customer_email,omitempty"`
	CustomerPhone   string `json:"customer_phone,omitempty"`
	VirtualAccount  string `json:"virtual_account,omitempty"`
	DateTimeExpired string `json:"datetime_expired,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SBNIRequest struct {
	ClientID string `json:"client_id,omitempty"`
	Data     string `json:"data,omitempty"`
}

//NewSBNIRequest : SBNIRequest Initializer
func NewSBNIRequest(DataStr string, BniClientID string) []byte {
	req := SBNIRequest{
		ClientID: BniClientID,
		Data:     DataStr,
	}
	reqJSONByte, err := json.Marshal(req)
	if err != nil {
		return []byte{}
	}
	return reqJSONByte
}

type BNIResponse struct {
	ClientID string `json:"client_id,omitempty"`
	Data     string `json:"data,omitempty"`
	Status   string `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
}
type SBNIResponseCreate struct {
	TrxID          string `json:"trx_id"`
	VirtualAccount string `json:"virtual_account"`
}

func RequestVABNI(booking_id string, total_charge string, fullName string, dueDate string) (string, error) {
	vaNumber := ""
	d := SBNICreateBilling{
		Type:         createBillingStr, // m
		ClientID:     bniClientID,      // m
		TrxID:        booking_id,       // m
		TrxAmount:    total_charge,     // m
		BillingType:  billingFixCode,   // m
		CustomerName: fullName,         // m
		// DateTimeExpired: dueDate.Format(time.RFC3339), // 1 H
		DateTimeExpired: dueDate,
		Description:     "RNA " + fullName,
	}
	mapB, err := json.Marshal(d)
	if err != nil {
		return vaNumber, errors.New("Error E8000 : Error marshal BNI request")
	}
	str := bniEnc.Encrypt(string(mapB), bniClientID, bniSecret)
	reqR := NewSBNIRequest(str, bniClientID)
	bodyBytes, _ := myHTTP.HTTPRequestJSON(trxURL, reqR)
	res := BNIResponse{}
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return vaNumber, errors.New("Error E8000 : Error marshal BNI request")
	}

	if res.Status != "000" {
		// a.Log.Error().Msgf("BookingOrder : BNI Response Failed %s", res.Status)
		return vaNumber, errors.New("Error E8001 : BNI Response Failed")
	}

	dt, _ := bniEnc.Decrypt(res.Data, bniClientID, bniSecret)
	rc := SBNIResponseCreate{}
	if err := json.Unmarshal([]byte(dt), &rc); err != nil {
		return vaNumber, errors.New("Error E8000 : Error marshal BNI request")
	}

	// TODO: add logger on bni response
	// a.Log.Info().Msgf("BookingOrder : BNI Response \n%s", res)
	vaNumber = rc.VirtualAccount
	return vaNumber, nil
}

//SBNIUpdateBilling : Update Billing info
type SBNIUpdateBilling struct {
	Type            string `json:"type,omitempty"`          // mandatory
	ClientID        string `json:"client_id,omitempty"`     // mandatory
	TrxID           string `json:"trx_id,omitempty"`        // mandatory
	TrxAmount       string `json:"trx_amount,omitempty"`    // mandatory
	CustomerName    string `json:"customer_name,omitempty"` // mandatory
	CustomerEmail   string `json:"customer_email,omitempty"`
	CustomerPhone   string `json:"customer_phone,omitempty"`
	DateTimeExpired string `json:"datetime_expired,omitempty"`
	Description     string `json:"description,omitempty"`
}

//BNIUpdateVA : Update VA billing
func DeactivateBNIInvoice(inv string, amount string, fullName string) error {
	d := SBNIUpdateBilling{
		Type:            updateBillingStr,
		ClientID:        bniClientID,
		TrxID:           inv,
		TrxAmount:       amount,
		CustomerName:    fullName,
		DateTimeExpired: helper.TimeToString(time.Now().Add(time.Minute * 2)),
		Description:     "Expired Now",
	}

	mapB, err := json.Marshal(d)
	if err != nil {
		return errors.New("Error E8000 : Error marshal BNI request")
	}
	str := bniEnc.Encrypt(string(mapB), bniClientID, bniSecret)
	reqR := NewSBNIRequest(str, bniClientID)

	bodyBytes, _ := myHTTP.HTTPRequestJSON(trxURL, reqR)
	res := BNIResponse{}
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return errors.New("Error E8000 : Error marshal BNI request")
	}

	if res.Status != "000" {
		// a.Log.Error().Msgf("BookingOrder : BNI Response Failed %s", res.Status)
		return errors.New("Error E8001 : BNI Response Failed")
	}

	dt, _ := bniEnc.Decrypt(res.Data, bniClientID, bniSecret)
	rc := SBNIResponseCreate{}
	if err := json.Unmarshal([]byte(dt), &rc); err != nil {
		return errors.New("Error E8000 : Error marshal BNI request")
	}
	return nil
}
