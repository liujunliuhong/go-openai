package completions

import (
	"configuration"
	"io"
	"net/http"
	"request"
)

// CompletionsRequest https://platform.openai.com/docs/api-reference/completions
type CompletionsRequest struct {
	c *configuration.Configuration
}

func NewCompletionsRequest(c *configuration.Configuration) *CompletionsRequest {
	return &CompletionsRequest{c: c}
}

func (c *CompletionsRequest) Method() string {
	return http.MethodPost
}

func (c *CompletionsRequest) Host() string {
	return request.BaseURL
}

func (c *CompletionsRequest) Path() string {
	return "/v1/models"
}

func (c *CompletionsRequest) URL() string {
	return c.Host() + c.Path()
}

func (c *CompletionsRequest) Headers() map[string]string {
	return c.c.Headers()
}

func (c *CompletionsRequest) Body() io.Reader {
	return nil
}
