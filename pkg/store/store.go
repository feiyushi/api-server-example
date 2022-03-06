package store

import (
	"apiserver/pkg/model"
	"strings"
	"sync"
)

// MetadataStore supports set metadata by id, get metadata by id and list metadata by company
type MetadataStore interface {
	SetMetadata(id string, metadata *model.Metadata) error
	GetMetadata(id string) (*model.Metadata, error)
	ListMedatadaByCompany(company string) ([]*model.Metadata, error)
}

// InMemoryMetadataStore implements MetadataStore interface and stores data in memory
type InMemoryMetadataStore struct {
	mutex    *sync.RWMutex              // supports concurrent metadata read/write
	metadata map[string]*model.Metadata // keyed by unique id
}

// ensures InMemoryMetadataStore implements MetadataStore interface
var _ MetadataStore = &InMemoryMetadataStore{}

// NewInMemoryMetadataStore creates a new InMemoryMetadataStore
func NewInMemoryMetadataStore() *InMemoryMetadataStore {
	return &InMemoryMetadataStore{
		mutex:    &sync.RWMutex{},
		metadata: make(map[string]*model.Metadata),
	}
}

// SetMetadata saves metadata into InMemoryMetadataStore
func (ms *InMemoryMetadataStore) SetMetadata(id string, metadata *model.Metadata) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.metadata[id] = metadata
	return nil
}

// GetMetadata gets metadata by id from InMemoryMetadataStore
// if not found, return nil data
func (ms *InMemoryMetadataStore) GetMetadata(id string) (*model.Metadata, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	if data, ok := ms.metadata[id]; ok {
		return data, nil
	}
	return nil, nil
}

// ListMedatadaByCompany list metadata that matches company name in InMemoryMetadataStore
// if not found, return empty slice
func (ms *InMemoryMetadataStore) ListMedatadaByCompany(company string) ([]*model.Metadata, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	var result []*model.Metadata
	for _, data := range ms.metadata {
		if strings.EqualFold(data.Company, company) {
			result = append(result, data)
		}
	}
	return result, nil
}
