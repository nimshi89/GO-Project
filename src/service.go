package src

import (
	"Go-Project/models"
)

type datasetStore interface {
	GetDatasets() models.Datasets
	GetDataset(id string) models.Dataset
}

type Service struct {
	Store datasetStore
}

// New creates a new service
func New(store_m datasetStore) *Service {
	svc := &Service{Store: store_m}
	return svc
}
