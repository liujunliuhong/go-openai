package completions

import (
	"bytes"
	"configuration"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"request"
)

// Request https://platform.openai.com/docs/api-reference/completions
type Request struct {
	c *configuration.Configuration
	p Param
}

func NewRequest(c *configuration.Configuration, p Param) *Request {
	return &Request{c: c, p: p}
}

func (c *Request) Method() string {
	return http.MethodPost
}

func (c *Request) Host() string {
	return request.BaseURL
}

func (c *Request) Path() string {
	return "/v1/completions"
}

func (c *Request) URL() string {
	return c.Host() + c.Path()
}

func (c *Request) Headers() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	for k, v := range c.c.Headers() {
		headers[k] = v
	}
	return headers
}

func (c *Request) Body() io.Reader {
	marshal, err := json.Marshal(c.p)
	if err != nil {
		return nil
	}
	fmt.Println(string(marshal))
	return bytes.NewBuffer(marshal)
}
