package store

import (
	"apiserver/pkg/model"
)

type MetadataStore interface {
	SetMetadata(id string, metadata *model.Metadata) error
	GetMetadata(id string) (*model.Metadata, error)
	ListMedatadaByCompany(company string) ([]*model.Metadata, error)
}
