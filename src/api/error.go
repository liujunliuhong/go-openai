package api

type APIError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Param   *string `json:"param,omitempty"`
	Type    *string `json:"type,omitempty"`
}
