package mp

import (
	"time"
	"github.com/davecgh/go-spew/spew"
)

const (
	// (GET) Looks for customers by many criteria.
	SearchCustomerUrl = "/v1/customers/search"

	// (POST) Makes a new customer.
	newCustomerPath = "/v1/customers/"

	// (GET) Retrieves information about a customer.
	GetCustomerUrl = "/v1/customers/:id"

	// (PUT) Updates a customer.
	SaveCustomerUrl = "/v1/customers/:id"

	// (DELETE) Removes a customer.
	RemoveCustomerUrl = "/v1/customers/:id"

	// (GET) Retrieves all cards from a customer.
	GetCustomerCardsUrl = "/v1/customers/:id/cards"

	// (POST) Saves a new card to the customer.
	NewCustomerCardUrl = "/v1/customers/:id/cards"

	// (GET) Retrieves information about a customer's card.
	GetCustomerCardUrl = "/v1/customers/:id/cards/:card_id"

	// (PUT) Updates a customer's card.
	SaveCustomerCardUrl = "/v1/customers/:id/cards/:card_id"

	// (DELETE) Removes a customer's card.
	RemoveCustomerCardUrl = "/v1/customers/:id/cards/:card_id"
)

func (c *Client) NewCustomer(customer *Customer) error {

	// create new customer
	err := c.Post(newCustomerPath, nil, nil, customer)

	spew.Dump(err, customer)

	return err
}

type Customer struct {
	CustomerId      string          `json:"id"`
	Email           string          `json:"email"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	Phone           *Phone          `json:"phone"`
	Identification  *Identification `json:"identification"`
	DefaultAddress  string          `json:"default_address"`
	Address         *MainAddress    `json:"address"`
	Description     string          `json:"description"`
	DateRegistered  time.Time       `json:"date_registered"`
	DateCreated     time.Time       `json:"date_created"`
	DateLastUpdated time.Time       `json:"date_last_updated"`
	Metadata        interface{}     `json:"metadata"`
	DefaultCard     string          `json:"default_card"`
	Cards           []Card          `json:"cards"`
	Addresses       []Address       `json:"addresses"`
	LiveMode        bool            `json:"live_mode"`
}

type Identification struct {
	Type    string `json:"type"`
	Number  string `json:"number"`
	Subtype string `json:"subtype"`
}

type MainAddress struct {
	AddressId    string `json:"id"`
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}

type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
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

type Address struct {
	AddressId     string         `json:"id"`
	Name          string         `json:"name"`
	Floor         string         `json:"floor"`
	Apartment     string         `json:"apartment"`
	StreetName    string         `json:"street_name"`
	StreetNumber  string         `json:"street_number"`
	ZipCode       string         `json:"zip_code"`
	City          *City          `json:"city"`
	State         *State         `json:"state"`
	Country       *Country       `json:"country"`
	Neighborhood  *Neighborhood  `json:"neighborhood"`
	Municipality  *Municipality  `json:"municipality"`
	Comments      string         `json:"comments"`
	DateCreated   time.Time      `json:"date_created"`
	Verifications *Verifications `json:"verifications"`
}

type City struct {
	CityId string `json:"id"`
	Name   string `json:"name"`
}

type State struct {
	StateId string `json:"id"`
	Name    string `json:"name"`
}

type Country struct {
	CountryId string `json:"id"`
	Name      string `json:"name"`
}

type Neighborhood struct {
	NeighborhoodId string `json:"id"`
	Name           string `json:"name"`
}

type Municipality struct {
	MunicipalityId string `json:"id"`
	Name           string `json:"name"`
}

type Verifications struct {
	Shipment     *Shipment `json:"shipment"`
	Municipality string    `json:"name"`
}

type Shipment struct {
	Success bool     `json:"success"`
	Errors  []Error  `json:"errors"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Field       string `json:"field"`
}
