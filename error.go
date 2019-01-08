package mp

import "encoding/json"

type ErrorImpl interface {
	GetMessage() string
	GetError() string
	GetStatus() int64
	HasCauses() bool
}

func (em ErrorMessage) GetMessage() string {
	return em.Message
}

func (em ErrorMessage) GetError() string {
	return em.Error
}

func (em ErrorMessage) GetStatus() int64 {
	status, _ := em.Status.Int64()
	return status
}

func (em ErrorMessage) HasCauses() bool {
	return len(em.Causes) > 0
}

type ErrorMessage struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Status  json.Number `json:"status"`
	Causes  []Cause     `json:"cause"`
}

type Cause struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Field       string `json:"field"`
}
