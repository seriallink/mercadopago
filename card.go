package mp

import "time"

type Card struct {
	CardId          string         `json:"id"`
	CustomerId      string         `json:"customer_id"`
	ExpirationMonth int            `json:"expiration_month"`
	ExpirationYear  int            `json:"expiration_year"`
	FirstSixDigits  string         `json:"first_six_digits"`
	LastFourDigits  string         `json:"last_four_digits"`
	PaymentMethod   *PaymentMethod `json:"payment_method"`
	SecurityCode    *SecurityCode  `json:"security_code"`
	Issuer          *Issuer        `json:"issuer"`
	CardHolder      *CardHolder    `json:"card_holder"`
	DateCreated     time.Time      `json:"date_created"`
	DateLastUpdated time.Time      `json:"date_last_updated"`
}

type PaymentMethod struct {
	PaymentId       string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeId   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

type SecurityCode struct {
	Length       int    `json:"length"`
	CardLocation string `json:"card_location"`
}

type Issuer struct {
	IssuerId string `json:"id"`
	Name     string `json:"name"`
}

type CardHolder struct {
	Name           string          `json:"name"`
	Identification *Identification `json:"identification"`
}