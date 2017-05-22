package mp

type ErrorImpl interface {
	GetMessage() string
	GetError() string
	GetStatus() int
	HasCauses() bool
}

func (em ErrorMessage) GetMessage() string {
	return em.Message
}

func (em ErrorMessage) GetError() string {
	return em.Error
}

func (em ErrorMessage) GetStatus() int {
	return em.Status
}

func (em ErrorMessage) HasCauses() bool {
	return len(em.Causes) > 0
}

type ErrorMessage struct {
	Message string  `json:"message"`
	Error   string  `json:"error"`
	Status  int     `json:"status"`
	Causes  []Cause `json:"cause"`
}

type Cause struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}