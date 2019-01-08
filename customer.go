package mp

import (
	"fmt"
)

const (
	getCustomerUrl        = "/v1/customers/"
	searchCustomerUrl     = "/v1/customers/search/"
	createCustomerPath    = "/v1/customers/"
	saveCustomerUrl       = "/v1/customers/"
	removeCustomerUrl     = "/v1/customers/"
	getCustomerCardsUrl   = "/v1/customers/%s/cards"
	createCustomerCardUrl = "/v1/customers/%s/cards"

	// (GET) Retrieves information about a customer's card.
	GetCustomerCardUrl = "/v1/customers/:id/cards/:card_id"

	// (PUT) Updates a customer's card.
	SaveCustomerCardUrl = "/v1/customers/:id/cards/:card_id"

	// (DELETE) Removes a customer's card.
	RemoveCustomerCardUrl = "/v1/customers/:id/cards/:card_id"
)

// Retrieves information about a customer.
func (c *Client) GetCustomer(customer *Customer) error {
	return c.Get(getCustomerUrl+customer.CustomerId, nil, nil, customer)
}

// Looks for customers by many criteria.
func (c *Client) SearchCustomer(customer *Customer) error {
	return c.Get(searchCustomerUrl, customer, nil, customer)
}

// Creates a new customer.
func (c *Client) CreateCustomer(customer *Customer) error {
	return c.Post(createCustomerPath, customer, nil, customer)
}

// Updates customer information.
func (c *Client) SaveCustomer(customer *Customer) error {
	return c.Put(saveCustomerUrl+customer.CustomerId, customer, nil, customer)
}

// Removes a customer.
func (c *Client) RemoveCustomer(customer *Customer) error {
	return c.Delete(removeCustomerUrl+customer.CustomerId, nil, nil, customer)
}

// Retrieves all cards from a customer.
func (c *Client) GetCustomerCards(customer *Customer) error {
	return c.Get(fmt.Sprintf(getCustomerCardsUrl, customer.CustomerId), nil, nil, customer)
}

// Saves a new card to the customer.
func (c *Client) NewCustomerCard(customer *Customer) error {
	return c.Post(fmt.Sprintf(createCustomerCardUrl, customer.CustomerId), customer, nil, customer)
}

type Customer struct {
	CustomerId      string            `json:"id,omitempty"`
	Email           string            `json:"email,omitempty"`
	FirstName       string            `json:"first_name,omitempty"`
	LastName        string            `json:"last_name,omitempty"`
	Phone           *Phone            `json:"phone,omitempty"`
	Identification  *Identification   `json:"identification,omitempty"`
	DefaultAddress  string            `json:"default_address,omitempty"`
	Address         *MainAddress      `json:"address,omitempty"`
	Description     string            `json:"description,omitempty"`
	DateRegistered  *Iso8601          `json:"date_registered"`
	DateCreated     *Iso8601          `json:"date_created,omitempty"`
	DateLastUpdated *Iso8601          `json:"date_last_updated,omitempty"`
	Metadata        interface{}       `json:"metadata,omitempty"`
	DefaultCard     string            `json:"default_card,omitempty"`
	Cards           []Card            `json:"cards,omitempty"`
	Addresses       []CustomerAddress `json:"addresses,omitempty"`
	LiveMode        bool              `json:"live_mode,omitempty"`
}

type MainAddress struct {
	AddressId    string `json:"id,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

type Phone struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

type CustomerAddress struct {
	AddressId     string         `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Floor         string         `json:"floor,omitempty"`
	Apartment     string         `json:"apartment,omitempty"`
	StreetName    string         `json:"street_name,omitempty"`
	StreetNumber  string         `json:"street_number,omitempty"`
	ZipCode       string         `json:"zip_code,omitempty"`
	City          *City          `json:"city,omitempty"`
	State         *State         `json:"state,omitempty"`
	Country       *Country       `json:"country,omitempty"`
	Neighborhood  *Neighborhood  `json:"neighborhood,omitempty"`
	Municipality  *Municipality  `json:"municipality,omitempty"`
	Comments      string         `json:"comments,omitempty"`
	DateCreated   *Iso8601       `json:"date_created,omitempty"`
	Verifications *Verifications `json:"verifications,omitempty"`
}

type City struct {
	CityId string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type State struct {
	StateId string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Country struct {
	CountryId string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Neighborhood struct {
	NeighborhoodId string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Municipality struct {
	MunicipalityId string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Verifications struct {
	Shipment     *Shipment `json:"shipment,omitempty"`
	Municipality string    `json:"name,omitempty"`
}

type Shipment struct {
	Success bool    `json:"success,omitempty"`
	Errors  []Cause `json:"errors,omitempty"`
}
