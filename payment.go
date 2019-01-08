package mp

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx/types"
)

const paymentUrl = "/v1/payments"

const (
	// Payment type
	RegularPayment    = "regular_payment"
	MoneyTransfer     = "money_transfer"
	RecurringPayment  = "recurring_payment"
	AccountFund       = "account_fund"
	PaymentAddition   = "payment_addition"
	CellphoneRecharge = "cellphone_recharge"
	PosPayment        = "pos_payment"

	// Type of order
	OrderTypeMercadoLibre = "mercadolibre"
	OrderTypeMercadoPago  = "mercadopago"

	// ID of the currency used in the payment
	ArgentinePeso     = "ARS"
	BrazilianReal     = "BRL"
	VenezuelanBolivar = "VEF"
	ChileanPeso       = "CLP"
	MexicanPeso       = "MXN"
	ColombianPeso     = "COP"
	PeruvianSol       = "PEN"
	UruguayanPeso     = "UYU"

	// Entity type of the payer (only for bank transfers)
	IndividualPayer  = "individual"
	AssociationPayer = "association"

	// Identification type of the associated payer (mandatory if the Payer is a Customer)
	CustomerType   = "customer"
	RegisteredType = "registered"
	GuestType      = "guest"

	// Payment status
	StatusPending     = "pending"
	StatusApproved    = "approved"
	StatusAuthorized  = "authorized"
	StatusInProcess   = "in_process"
	StatusInMediation = "in_mediation"
	StatusRejected    = "rejected"
	StatusCancelled   = "cancelled"
	StatusRefunded    = "refunded"
	StatusChargedBack = "charged_back"

	// Fee detail
	MercadoPagoFee = "mercadopago_fee"
	CouponFee      = "coupon_fee"
	FinancingFee   = "financing_fee"
	ShippingFee    = "shipping_fee"
	ApplicationFee = "application_fee"
	DiscountFee    = "discount_fee"

	// Who absorbs the cost
	FeeCollector = "collector"
	FeePayer     = "payer"

	// Type of payment method chosen
	AccountMoney = "account_money"
	Ticket       = "ticket"
	BankTransfer = "bank_transfer"
	ATM          = "atm"
	CreditCard   = "credit_card"
	DebitCard    = "debit_card"
	PrepaidCard  = "prepaid_card"

	// Type of user who issued the refund
	RefundCollector = "collector"
	RefundOperator  = "operator"
	RefundAdmin     = "admin"
	RefundBPP       = "bpp"

	// Encoding type
	UCCEAN128 = "UCC/EAN 128"
	Code128C  = "Code128C"
	Code39    = "Code39"
)

type Payment struct {
	Id                        int64               `json:"id,omitempty"`
	DateCreated               *Iso8601            `json:"date_created,omitempty"`
	DateApproved              *Iso8601            `json:"date_approved,omitempty"`
	DateOfExpiration          *Iso8601            `json:"date_of_expiration,omitempty"`
	DateLastUpdated           *Iso8601            `json:"date_last_updated,omitempty"`
	MoneyReleaseDate          *Iso8601            `json:"money_release_date,omitempty"`
	CollectorId               int64               `json:"collector_id,omitempty"`
	OperationType             string              `json:"operation_type,omitempty"`
	Payer                     *Payer              `json:"payer,omitempty"`
	BinaryMode                bool                `json:"binary_mode,omitempty"`
	LiveMode                  bool                `json:"live_mode,omitempty"`
	Order                     *Order              `json:"order,omitempty"`
	ExternalReference         string              `json:"external_reference,omitempty"`
	Description               string              `json:"description,omitempty"`
	Metadata                  types.JSONText      `json:"metadata"`
	CurrencyId                string              `json:"currency_id,omitempty"`
	TransactionAmount         float64             `json:"transaction_amount,omitempty"`
	TransactionAmountRefunded float64             `json:"transaction_amount_refunded,omitempty"`
	CouponAmount              float64             `json:"coupon_amount,omitempty"`
	CampaignId                int64               `json:"campaign_id,omitempty"`
	CouponCode                string              `json:"coupon_code,omitempty"`
	TransactionDetails        *TransactionDetails `json:"transaction_details"`
	FeeDetails                []Fee               `json:"fee_details,omitempty"`
	DifferentialPricingId     int64               `json:"differential_pricing_id,omitempty"`
	ApplicationFee            float64             `json:"application_fee,omitempty"`
	Status                    json.Number         `json:"status,omitempty"`
	StatusDetail              string              `json:"status_detail,omitempty"`
	Capture                   bool                `json:"capture,omitempty"`
	CallForAuthorizeId        string              `json:"call_for_authorize_id,omitempty"`
	PaymentMethodId           string              `json:"payment_method_id,omitempty"`
	IssuerId                  string              `json:"issuer_id,omitempty"`
	PaymentTypeId             string              `json:"payment_type_id,omitempty"`
	Token                     string              `json:"token,omitempty"`
	Card                      *Card               `json:"card,omitempty"`
	StatementDescriptor       string              `json:"statement_descriptor,omitempty"`
	Installments              int64               `json:"installments,omitempty"`
	NotificationUrl           string              `json:"notification_url,omitempty"`
	CallbackUrl               string              `json:"callback_url,omitempty"`
	Refunds                   []Refund            `json:"refunds,omitempty"`
	AdditionalInfo            *AdditionalInfo     `json:"additional_info"`
}

type Payer struct {
	Id             string          `json:"id,omitempty"`
	Email          string          `json:"email,omitempty"`
	FirstName      string          `json:"first_name,omitempty"`
	LastName       string          `json:"last_name,omitempty"`
	EntityType     string          `json:"entity_type,omitempty"`
	Type           string          `json:"type,omitempty"`
	Identification *Identification `json:"identification"`
	Address        *Address        `json:"address,omitempty"`
	Phone          *Phone          `json:"phone,omitempty"`
}

type Address struct {
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"federal_unit,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	Country      string `json:"country,omitempty"`
}

type Order struct {
	Id   int64  `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type TransactionDetails struct {
	FinancialInstitution     string  `json:"financial_institution,omitempty"`
	NetReceivedAmount        float64 `json:"net_received_amount,omitempty"`
	TotalPaidAmount          float64 `json:"total_paid_amount,omitempty"`
	InstallmentAmount        float64 `json:"installment_amount,omitempty"`
	OverpaidAmount           float64 `json:"overpaid_amount,omitempty"`
	ExternalResourceUrl      string  `json:"external_resource_url,omitempty"`
	PaymentMethodReferenceId string  `json:"payment_method_reference_id,omitempty"`
}

type Fee struct {
	Type     string  `json:"type,omitempty"`
	FeePayer string  `json:"fee_payer,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
}

type Refund struct {
	Id                   string         `json:"id,omitempty"`
	PaymentId            int64          `json:"payment_id,omitempty"`
	Amount               float64        `json:"amount,omitempty"`
	Metadata             types.JSONText `json:"metadata,omitempty"`
	Source               *Source        `json:"source,omitempty"`
	DateCreated          *Iso8601       `json:"date_created,omitempty"`
	UniqueSequenceNumber string         `json:"unique_sequence_number,omitempty"`
}

type Source struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Data that could improve fraud analysis and conversion rates. Try to send as much information as possible.
type AdditionalInfo struct {
	IpAddress string     `json:"ip_address,omitempty"`
	Items     []Item     `json:"items,omitempty"`
	Buyer     *Buyer     `json:"payer,omitempty"`
	Shipments *Shipments `json:"shipments,omitempty"`
	Barcode   *Barcode   `json:"barcode,omitempty"`
}

type Item struct {
	Id          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureUrl  string  `json:"picture_url,omitempty"`
	CategoryId  string  `json:"category_id,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
}

type Buyer struct {
	FirstName        string   `json:"first_name,omitempty"`
	LastName         string   `json:"last_name,omitempty"`
	Phone            *Phone   `json:"phone,omitempty"`
	Address          *Address `json:"address,omitempty"`
	RegistrationDate *Iso8601 `json:"registration_date,omitempty"`
}

type Shipments struct {
	ReceiverAddress *ReceiverAddress `json:"receiver_address,omitempty"`
}

type ReceiverAddress struct {
	ZipCode      string      `json:"zip_code,omitempty"`
	StreetName   string      `json:"street_name,omitempty"`
	StreetNumber json.Number `json:"street_number,omitempty"`
	Floor        string      `json:"floor,omitempty"`
	Apartment    string      `json:"apartment,omitempty"`
}

type Barcode struct {
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	Width   int64  `json:"width,omitempty"`
	Height  int64  `json:"height,omitempty"`
}

// Retrieves information about a payment
func (c *Client) GetPayment(payment *Payment) error {
	return c.Get(fmt.Sprintf("%s/%d", paymentUrl, payment.Id), nil, nil, payment)
}

// Issues a new payment
func (c *Client) CreatePayment(payment *Payment) error {
	return c.Post(paymentUrl, payment, nil, payment)
}

// Updates a payment
func (c *Client) UpdatePayment(payment *Payment) error {
	return c.Put(fmt.Sprintf("%s/%d", paymentUrl, payment.Id), payment, nil, payment)
}
