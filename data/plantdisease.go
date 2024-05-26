package data

import (
	"time"
)

type PlantDisease struct {
	Disease     *Disease               `json:"disease"`
	Plant       *Plant                 `json:"plant"`
	Description map[string]interface{} `json:"description"`
	CreatedAt   time.Time              `json:"created_at"`
}

type CreatePlantDiseaseRequest struct {
	PlantName   string                 `json:"plant_name"`
	DiseaseName string                 `json:"disease_name"`
	Description map[string]interface{} `json:"description"`
}

type CreatePlantDiseasesRequest struct {
	Diseases []*CreatePlantDiseaseRequest `json:"plantdiseases"`
}

func NewPlantDisease(plantdisease *CreatePlantDiseaseRequest) (*PlantDisease, error) {
	plant, err := NewPlant(&CreatePlantRequest{
		PlantName: plantdisease.PlantName,
	})
	if err != nil {
		return nil, err
	}

	disease, err := NewDisease(&CreateDiseaseRequest{
		DiseaseName: plantdisease.DiseaseName,
	})
	if err != nil {
		return nil, err
	}

	return &PlantDisease{
		Plant:       plant,
		Disease:     disease,
		Description: plantdisease.Description,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
