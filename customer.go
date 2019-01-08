package mp

import (
	"fmt"
	"time"
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
	CustomerId      string          `json:"id,omitempty"`
	Email           string          `json:"email,omitempty"`
	FirstName       string          `json:"first_name,omitempty"`
	LastName        string          `json:"last_name,omitempty"`
	Phone           *Phone          `json:"phone"`
	Identification  *Identification `json:"identification"`
	DefaultAddress  string          `json:"default_address,omitempty"`
	Address         *MainAddress    `json:"address"`
	Description     string          `json:"description,omitempty"`
	DateRegistered  *time.Time      `json:"date_registered"`
	DateCreated     *time.Time      `json:"date_created"`
	DateLastUpdated *time.Time      `json:"date_last_updated"`
	Metadata        interface{}     `json:"metadata"`
	DefaultCard     string          `json:"default_card,omitempty"`
	Cards           []Card          `json:"cards"`
	Addresses       []Address       `json:"addresses"`
	LiveMode        bool            `json:"live_mode"`
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
	Success bool    `json:"success"`
	Errors  []Cause `json:"errors"`
}
