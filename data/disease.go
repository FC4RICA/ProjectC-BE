package data

import (
	"time"
)

type Disease struct {
	ID          int                    `json:"id"`
	DiseaseName string                 `json:"disease_name"`
	PlantName   string                 `json:"plant_name"`
	Description map[string]interface{} `json:"description"`
	CreatedAt   time.Time              `json:"created_at"`
}

type CreateDiseaseRequest struct {
	DiseaseName string                 `json:"disease_name"`
	PlantName   string                 `json:"plant_name"`
	Description map[string]interface{} `json:"description"`
}

func NewDisease(disease *CreateDiseaseRequest) (*Disease, error) {
	return &Disease{
		DiseaseName: disease.DiseaseName,
		PlantName:   disease.PlantName,
		Description: disease.Description,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
