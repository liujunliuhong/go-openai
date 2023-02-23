package models

type Model struct {
	Id         string       `json:"id"`
	Object     string       `json:"object"`
	Created    int64        `json:"created"`
	OwnedBy    string       `json:"owned_by"`
	Root       string       `json:"root"`
	Parent     string       `json:"parent"`
	Permission []Permission `json:"permission"`
}

type ModeList struct {
	Models []Model `json:"data"`
}

type Permission struct {
}
