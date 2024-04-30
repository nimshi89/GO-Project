package src

import "project/models"

type datasetStore interface {
	get_datasets() models.Dataset
	get_datset(id string) models.Datasets
}

type Service struct {
	store *datasetStore
}

// New creates a new service
func New(store_m *datasetStore) *Service {
	svc := &Service{
		store: store_m,
	}
	return svc
}
