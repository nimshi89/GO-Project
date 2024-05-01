package stores

import "Go-Project/models"

type Stub struct {
	Path string
}

func (stub *Stub) GetDataset(id string) models.Dataset {
	dataset := models.Dataset{Id: "cpih", PublicationDate: "24/4/2024T00:00:00"}
	return dataset
}

func (stub *Stub) GetDatasets() models.Datasets {
	return models.Datasets{}
}
