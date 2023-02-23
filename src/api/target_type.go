package api

import (
	"request"
	"strings"
)

type RequestType int

const baseURL = "https://api.openai.com"

const (
	ListModels RequestType = iota
	Completion RequestType = iota
)

func (r RequestType) Path() string {
	switch r {
	case ListModels:
		return "/v1/models"
	case Completion:
		return "/v1/completions"
	}
	return ""
}

func (r RequestType) Method() string {
	switch r {
	case ListModels:
		return "GET"
	case Completion:
		return "POST"
	}
	return "POST"
}

func (r RequestType) Header(config request.Configuration) map[string]string {
	return map[string]string{
		"": "",
	}
}

func (r RequestType) URL() string {
	if !strings.HasSuffix(r.Path(), "/") {
		return baseURL + "/" + r.Path()
	}
	return baseURL + r.Path()
}
