package data

import (
	"time"
)

type CreateResultRequest struct {
	UserID int      `json:"user_id"`
	Images []string `json:"images"`
}

type Result struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	DiseaseID     int       `json:"disease_id"`
	PredictResult bool      `json:"predict_result"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewResult(result *CreateResultRequest) (*Result, error) {
	// call ai api
	diseaseID := 0
	predictResult := false
	return &Result{
		UserID:        result.UserID,
		DiseaseID:     diseaseID,
		PredictResult: predictResult,
		CreatedAt:     time.Now().UTC(),
	}, nil
}
