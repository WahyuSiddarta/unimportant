package structure

import "encoding/json"

type SPaymentMethod struct {
	PaymentID         int    `db:"payment_method_id"`
	PaymentTypeName   string `db:"payment_type_name"`
	PaymentMethodName string `db:"payment_method_name"`
	PaymentType       int    `db:"payment_type"`
	AccNumber         string `db:"acc_number"`
	AccName           string `db:"acc_name"`
	BankName          string `db:"bank_name"`
	BankLogo          string `db:"bank_logo"`
}

func (a SPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PaymentID         int    `json:"payment_method_id"`
		PaymentTypeName   string `json:"payment_type_name"`
		PaymentMethodName string `json:"payment_method_name"`
		PaymentType       int    `json:"payment_type"`
		AccNumber         string `json:"acc_number"`
		AccName           string `json:"acc_name"`
		BankName          string `json:"bank_name"`
		BankLogo          string `json:"bank_logo"`
	}{
		PaymentID:         a.PaymentID,
		PaymentTypeName:   a.PaymentTypeName,
		PaymentMethodName: a.PaymentMethodName,
		PaymentType:       a.PaymentType,
		AccNumber:         a.AccNumber,
		AccName:           a.AccName,
		BankName:          a.BankName,
		BankLogo:          a.BankLogo,
	})
}
