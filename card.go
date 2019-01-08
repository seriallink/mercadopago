package mp

import (
	"github.com/davecgh/go-spew/spew"
	"time"
)

const (
	cardTokenUrl = "/v1/card_tokens/"
)

// Creates a new card token.
func (c *Client) CreateCardToken(cardtoken *CardToken) error {
	err := c.Post(cardTokenUrl, cardtoken, nil, cardtoken)
	spew.Dump(err, cardtoken)
	return err
}

// Retrieves information about a card token.
func (c *Client) GetCardToken(cardtoken *CardToken) error {
	return c.Get(cardTokenUrl+cardtoken.CardTokenId, nil, nil, cardtoken)
}

// Updates an existing card token.
func (c *Client) SaveCardToken(cardtoken *CardToken) error {
	return c.Put(cardTokenUrl+cardtoken.CardTokenId, cardtoken, nil, cardtoken)
}

type CardToken struct {
	CardTokenId        string      `json:"id,omitempty"`
	CardId             string      `json:"cardId,omitempty"`
	PublicKey          string      `json:"publicKey,omitempty"`
	LuhnValidation     bool        `json:"luhnValidation"`
	FirstSixDigits     string      `json:"firstSixDigits,omitempty"`
	LastFourDigits     string      `json:"lastFourDigits,omitempty"`
	Status             string      `json:"status,omitempty"`
	CardNumberLength   int         `json:"cardNumberLength"`
	SecurityCodeLength int         `json:"securityCodeLength"`
	ExpirationYear     int         `json:"expirationYear"`
	ExpirationMonth    int         `json:"expirationMonth"`
	CardHolder         *CardHolder `json:"cardholder"`
	DateDue            *time.Time  `json:"dateDue"`
	DateUsed           *time.Time  `json:"dateUsed"`
	DateCreated        *time.Time  `json:"dateCreated"`
	DateLastUpdated    *time.Time  `json:"dateLastUpdated"`
	LiveMode           bool        `json:"liveMode"`
}

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

type SecurityCode struct {
	Mode         string `json:"mode"`
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

type CardNumber struct {
	Length     int    `json:"length"`
	Validation string `json:"validation"`
}
