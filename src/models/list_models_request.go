package models

import (
	"configuration"
	"io"
	"net/http"
	"request"
)

// ListModelsRequest https://platform.openai.com/docs/api-reference/models/list
type ListModelsRequest struct {
	c *configuration.Configuration
}

func NewListModelsRequest(c *configuration.Configuration) *ListModelsRequest {
	return &ListModelsRequest{c: c}
}

func (l *ListModelsRequest) Method() string {
	return http.MethodGet
}

func (l *ListModelsRequest) Host() string {
	return request.BaseURL
}

func (l *ListModelsRequest) Path() string {
	return "/v1/models"
}

func (l *ListModelsRequest) URL() string {
	return l.Host() + l.Path()
}

func (l *ListModelsRequest) Headers() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	for k, v := range l.c.Headers() {
		headers[k] = v
	}
	return headers
}

func (l *ListModelsRequest) Body() io.Reader {
	return nil
}
