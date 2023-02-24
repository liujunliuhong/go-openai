package models

import (
	"configuration"
	"request_handler"
)

func List(c *configuration.Configuration) (response Response, err error) {
	r := NewListModelsRequest(c)
	err = request_handler.Perform(r, &response)
	return
}

func Retrieve(c *configuration.Configuration, id string) (model Model, err error) {
	r := NewRetrieveModelRequest(c, id)
	err = request_handler.Perform(r, &model)
	return
}
