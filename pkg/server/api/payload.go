package api

import "apiserver/pkg/model"

// API payload definition.
// For simplicity, it's the same as core model (plus unique id).
// Otherwise as API can change over time, conversion is needed between payload contract and core model
type Payload struct {
	ID   string          `yaml:"id"`
	Data *model.Metadata `yaml:"metadata"`
}

func NewPayload(id string, data *model.Metadata) *Payload {
	return &Payload{
		ID:   id,
		Data: data,
	}
}
