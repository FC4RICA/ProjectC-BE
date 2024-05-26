package data

import "time"

type Disease struct {
	ID          int       `json:"id"`
	DiseaseName string    `json:"disease_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateDiseaseRequest struct {
	DiseaseName string `json:"disease_name"`
}

func NewDisease(disease *CreateDiseaseRequest) (*Disease, error) {
	return &Disease{
		DiseaseName: disease.DiseaseName,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
