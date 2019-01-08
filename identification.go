package mp

const identificationTypesUrl = "/v1/identification_types"

type Identification struct {
	Id        string `json:"id,omitempty"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	Number    string `json:"number,omitempty"`
	Subtype   string `json:"subtype,omitempty"`
	MinLength int    `json:"min_length,omitempty"`
	MaxLength int    `json:"max_length,omitempty"`
}

// Returns information about identification types.
func (c *Client) GetIdentificationTypes() (ids []Identification, err error) {
	err = c.Get(identificationTypesUrl, nil, nil, &ids)
	return
}
