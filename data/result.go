package data

import (
	"time"
)

type CreateResultRequest struct {
	UserID       int           `json:"user_id"`
	PlantDisease *PlantDisease `json:"disease_id"`
}

type Result struct {
	ID           int           `json:"id"`
	UserID       int           `json:"user_id"`
	PlantDisease *PlantDisease `json:"disease_id"`
	CreatedAt    time.Time     `json:"created_at"`
}

type PredictResponse struct {
	PlantName   string `json:"plantname"`
	DiseaseName string `json:"diseasename"`
}

func NewResult(result *CreateResultRequest) (*Result, error) {
	return &Result{
		UserID:       result.UserID,
		PlantDisease: result.PlantDisease,
		CreatedAt:    time.Now().UTC(),
	}, nil
}
