package mp

// https://www.mercadopago.com.br/developers/en/reference/preferences/resource/
const (
	getPreferenceUrl    = "/checkout/preferences/"
	createPreferenceUrl = "/checkout/preferences/"
	savePreferenceUrl   = "/checkout/preferences/"
)

type Preference struct {
	Items               []Item               `json:"items,omitempty"`
	Payer               *Payer               `json:"payer,omitempty"`
	PaymentMethods      *PaymentMethods      `json:"payment_methods,omitempty"`
	Shipments           *PreferenceShipments `json:"shipments,omitempty"`
	BackUrls            map[string]string    `json:"back_urls,omitempty"`
	NotificationUrl     string               `json:"notification_url,omitempty"`
	PreferenceId        string               `json:"id,omitempty"`
	InitPoint           string               `json:"init_point,omitempty"`
	SandboxInitPoint    string               `json:"sandbox_init_point,omitempty"`
	DataCreated         *Iso8601             `json:"data_created,omitempty"`
	OperationType       string               `json:"operation_type,omitempty"`
	AdditionalInfo      string               `json:"additional_info,omitempty"`
	AutoReturn          string               `json:"auto_return,omitempty"`
	ExternalReference   string               `json:"external_reference,omitempty"`
	Expires             bool                 `json:"expires,omitempty"`
	ExpirationDateFrom  *Iso8601             `json:"expiration_date_from,omitempty"`
	ExpirationDateTo    *Iso8601             `json:"expiration_date_to,omitempty"`
	CollectorId         int64                `json:"collector_id,omitempty"`
	ClientId            string               `json:"client_id,omitempty"`
	Marketplace         string               `json:"marketplace,omitempty"`
	MarketplaceFee      float64              `json:"marketplace_fee,omitempty"`
	DifferentialPricing *DifferentialPricing `json:"differential_pricing,omitempty"`
	BinaryMode          bool                 `json:"binary_mode,omitempty"`
	Taxes               []Taxes              `json:"taxes,omitempty"`
}

type DifferentialPricing struct {
	Id int `json:"id,omitempty"`
}

type PaymentMethods struct {
	Id                     string                  `json:"id,omitempty"`
	ExcludedPaymentMethods []ExcludedPaymentMethod `json:"excluded_payment_method,omitempty"`
	ExcludedPaymentTypes   []ExcludedPaymentTypes  `json:"excluded_payment_types,omitempty"`
	DefaultPaymentMethodId string                  `json:"default_payment_method_id,omitempty"`
	Installments           int64                   `json:"installments,omitempty"`
	DefaultInstallments    int64                   `json:"defaultInstallments,omitempty"`
}

type ExcludedPaymentMethod struct {
	Id string `json:"id,omitempty"`
}

type ExcludedPaymentTypes struct {
	Id string `json:"id,omitempty"`
}

type PreferenceShipments struct {
	Mode                  string           `json:"mode,omitempty"`
	LocalPickup           bool             `json:"local_pickup,omitempty"`
	Dimensions            string           `json:"dimensions,omitempty"`
	DefaultShippingMethod string           `json:"default_shipping_method,omitempty"`
	FreeMethods           []ShippingMethod `json:"free_methods,omitempty"`
	Cost                  float64          `json:"id,omitempty"`
	FreeShipping          bool             `json:"free_shipping,omitempty"`
	ReceiverAddress       *ReceiverAddress `json:"receiver_address,omitempty"`
}

type ShippingMethod struct {
	Id string `json:"id,omitempty"`
}

type Taxes struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// Retrieves information about a preference
func (c *Client) GetPreference(preference *Preference) error {
	return c.Get(getPreferenceUrl+preference.PreferenceId, nil, nil, preference)
}

// Creates a new preference.
func (c *Client) CreatePreference(preference *Preference) error {
	return c.Post(createPreferenceUrl, preference, nil, preference)
}

// Updates preference information.
func (c *Client) SavePreference(preference *Preference) error {
	return c.Put(savePreferenceUrl+preference.PreferenceId, preference, nil, preference)
}
