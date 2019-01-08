package mp

const paymentMethodsUrl = "/v1/payment_methods"

// Retrieves information about available payment methods.
func (c *Client) GetPaymentMethods() (pms []PaymentMethod, err error) {
	err = c.Get(paymentMethodsUrl, nil, nil, &pms)
	return
}

type PaymentMethod struct {
	PaymentId             string                 `json:"id"`
	Name                  string                 `json:"name"`
	PaymentTypeId         string                 `json:"payment_type_id"`
	Status                string                 `json:"status"`
	SecureThumbnail       string                 `json:"secure_thumbnail"`
	Thumbnail             string                 `json:"thumbnail"`
	DeferredCapture       string                 `json:"deferred_capture"`
	AdditionalInfoNeeded  []string               `json:"additional_info_needed"`
	MinAllowedAmount      float64                `json:"min_allowed_amount"`
	MaxAllowedAmount      float64                `json:"max_allowed_amount"`
	AccreditationTime     int64                  `json:"accreditation_time"`
	Settings              []Settings             `json:"settings"`
	FinancialInstitutions []FinancialInstitution `json:"financial_institutions"`
}

type Settings struct {
	Bin          *Bin          `json:"bin"`
	CardNumber   *CardNumber   `json:"card_number"`
	SecurityCode *SecurityCode `json:"security_code"`
}

type Bin struct {
	Pattern             string `json:"pattern"`
	ExclusionPattern    string `json:"exclusion_pattern"`
	InstallmentsPattern string `json:"installments_pattern"`
}

type FinancialInstitution struct {
	FinancialInstitutionId string `json:"id"`
	Description            string `json:"description"`
}
