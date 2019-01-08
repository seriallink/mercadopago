package mp

const (
	cardTokenUrl = "/v1/card_tokens/"
)

// Creates a new card token.
func (c *Client) CreateCardToken(cardtoken *CardToken) error {
	return c.Post(cardTokenUrl, cardtoken, nil, cardtoken)
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
	CardNumberLength   int         `json:"cardNumberLength,omitempty"`
	SecurityCodeLength int         `json:"securityCodeLength,omitempty"`
	ExpirationYear     int         `json:"expirationYear,omitempty"`
	ExpirationMonth    int         `json:"expirationMonth,omitempty"`
	CardHolder         *CardHolder `json:"cardholder,omitempty"`
	DateDue            *Iso8601    `json:"dateDue,omitempty"`
	DateUsed           *Iso8601    `json:"dateUsed,omitempty"`
	DateCreated        *Iso8601    `json:"dateCreated,omitempty"`
	DateLastUpdated    *Iso8601    `json:"dateLastUpdated,omitempty"`
	LiveMode           bool        `json:"liveMode,omitempty"`
}

type Card struct {
	CardId          string         `json:"id,omitempty"`
	CustomerId      string         `json:"customer_id,omitempty"`
	ExpirationMonth int            `json:"expiration_month,omitempty"`
	ExpirationYear  int            `json:"expiration_year,omitempty"`
	FirstSixDigits  string         `json:"first_six_digits,omitempty"`
	LastFourDigits  string         `json:"last_four_digits,omitempty"`
	PaymentMethod   *PaymentMethod `json:"payment_method,omitempty"`
	SecurityCode    *SecurityCode  `json:"security_code,omitempty"`
	Issuer          *Issuer        `json:"issuer,omitempty"`
	CardHolder      *CardHolder    `json:"card_holder,omitempty"`
	DateCreated     *Iso8601       `json:"date_created,omitempty"`
	DateLastUpdated *Iso8601       `json:"date_last_updated,omitempty"`
}

type SecurityCode struct {
	Mode         string `json:"mode,omitempty"`
	Length       int    `json:"length,omitempty"`
	CardLocation string `json:"card_location,omitempty"`
}

type Issuer struct {
	IssuerId string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
}

type CardHolder struct {
	Name           string          `json:"name,omitempty"`
	Identification *Identification `json:"identification,omitempty"`
}

type CardNumber struct {
	Length     int    `json:"length,omitempty"`
	Validation string `json:"validation,omitempty"`
}
