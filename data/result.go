package data

import (
	"time"
)

type CreateResultRequest struct {
	UserID        int  `json:"user_id"`
	DiseaseID     int  `json:"disease_id"`
	PredictResult bool `json:"predict_result"`
}

type Result struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	DiseaseID     int       `json:"disease_id"`
	PredictResult bool      `json:"predict_result"`
	CreatedAt     time.Time `json:"created_at"`
}

type PredictResponse struct {
	PlantName   string `json:"plantname"`
	DiseaseName string `json:"diseasename"`
}

func NewResult(result *CreateResultRequest) (*Result, error) {
	return &Result{
		UserID:        result.UserID,
		DiseaseID:     result.DiseaseID,
		PredictResult: result.PredictResult,
		CreatedAt:     time.Now().UTC(),
	}, nil
}
