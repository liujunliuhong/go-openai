package models

import (
	"request"
	"request_handler"
)

func List(c *request.Configuration) (models ModeList, err error) {
	r := NewListModelsRequest(c)
	err = request_handler.Perform(r, &models)
	return
}
