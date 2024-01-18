package models

type PaymentRequest struct {
	PaymentType string `json:"payment_type"`
	GrossAmount int    `json:"gross_amount"`
}
