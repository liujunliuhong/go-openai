package request

import "io"

const BaseURL = "https://api.openai.com"

type Request interface {
	Host() string
	Path() string
	URL() string
	Method() string
	Headers() map[string]string
	Body() io.Reader
}
