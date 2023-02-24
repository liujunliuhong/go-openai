package completions

import (
	"configuration"
	"request_handler"
)

func Create(c *configuration.Configuration, p Param) (response Response, err error) {
	r := NewRequest(c, p)
	err = request_handler.Perform(r, &response)
	return
}
