package mp

import "encoding/json"

type ErrorImpl interface {
	Error() string
	GetMessage() string
	GetDetail() string
	GetStatus() int64
	HasCauses() bool
}

func (em ErrorMessage) Error() string {
	if em.HasCauses() {
		return em.Causes[0].Description
	}
	return em.GetMessage()
}

func (em ErrorMessage) GetMessage() string {
	return em.Message
}

func (em ErrorMessage) GetDetail() string {
	return em.Detail
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
	Detail  string      `json:"error"`
	Status  json.Number `json:"status"`
	Causes  []Cause     `json:"cause"`
}

type Cause struct {
	Code        json.Number `json:"code"`
	Description string      `json:"description"`
	Field       string      `json:"field"`
}
