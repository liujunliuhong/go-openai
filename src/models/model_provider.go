package models

import (
	"configuration"
	"request_handler"
)

func List(c *configuration.Configuration) (models ModeList, err error) {
	r := NewListModelsRequest(c)
	err = request_handler.Perform(r, &models)
	return
}

func Retrieve(c *configuration.Configuration, id string) (model Model, err error) {
	r := NewRetrieveModelRequest(c, id)
	err = request_handler.Perform(r, &model)
	return
}
