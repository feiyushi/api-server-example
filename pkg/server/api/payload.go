package api

import "apiserver/pkg/model"

// API payload definition.
// For simplicity, it's the same as core model.
// Otherwise as API can change over time, conversion is needed between payload contract and core model
type Payload = model.MetadataWithID

func NewPayload(id string, data *model.Metadata) *Payload {
	return &Payload{
		ID:   id,
		Data: data,
	}
}
