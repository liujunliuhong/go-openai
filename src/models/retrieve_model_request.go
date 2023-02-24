package models

import (
	"configuration"
	"fmt"
	"io"
	"net/http"
	"request"
)

// RetrieveModelRequest https://platform.openai.com/docs/api-reference/models/retrieve
type RetrieveModelRequest struct {
	c       *configuration.Configuration
	modelID string
}

func NewRetrieveModelRequest(c *configuration.Configuration, modelID string) *RetrieveModelRequest {
	return &RetrieveModelRequest{c: c, modelID: modelID}
}

func (r *RetrieveModelRequest) Method() string {
	return http.MethodGet
}

func (r *RetrieveModelRequest) Host() string {
	return request.BaseURL
}

func (r *RetrieveModelRequest) Path() string {
	return fmt.Sprintf("/v1/models/%s", r.modelID)
}

func (r *RetrieveModelRequest) URL() string {
	return r.Host() + r.Path()
}

func (r *RetrieveModelRequest) Headers() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	for k, v := range r.c.Headers() {
		headers[k] = v
	}
	return headers
}

func (r *RetrieveModelRequest) Body() io.Reader {
	return nil
}
