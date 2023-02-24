package openai_error

import "fmt"

type APIError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Param   *string `json:"param,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type APIErrorResponse struct {
	Err *APIError `json:"error,omitempty"`
}

func (a *APIErrorResponse) Error() string {
	return fmt.Sprintf("Status Code: %s, Message: %s", *(a.Err.Code), *(a.Err.Message))
}

type RequestError struct {
	StatusCode int
	Err        error
}

func (r RequestError) Error() string {
	return fmt.Sprintf("Status Code: %d, Message: %s", r.StatusCode, r.Err.Error())
}
