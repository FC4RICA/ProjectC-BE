package db

import (
	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateResult(result *data.Result) (int, error) {
	query := `INSERT INTO Result
		(user_id, disease_id, predict_result, created_at)
		VALUES ($1, $2, $3, $4) RETURNING result_id`

	var id int
	err := s.db.QueryRow(
		query,
		result.UserID, result.DiseaseID, result.PredictResult, result.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
