package db

import (
	"database/sql"
	"fmt"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateResult(result *data.Result) (int, error) {
	query := `INSERT INTO PredictResult
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

func (s *PostgresStore) GetResultsByUserID(id int) ([]*data.Result, error) {
	rows, err := s.db.Query("SELECT * FROM PredictResult WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	results := []*data.Result{}
	for rows.Next() {
		result, err := scanIntoResult(rows)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (s *PostgresStore) GetResultByID(id int) (*data.Result, error) {
	rows, err := s.db.Query("SELECT * FROM PredictResult WHERE result_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoResult(rows)
	}

	return nil, fmt.Errorf("result %d not found", id)
}

func (s *PostgresStore) GetResults() ([]*data.Result, error) {
	rows, err := s.db.Query("SELECT * FROM PredictResult")
	if err != nil {
		return nil, err
	}

	results := []*data.Result{}
	for rows.Next() {
		result, err := scanIntoResult(rows)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func scanIntoResult(rows *sql.Rows) (*data.Result, error) {
	result := new(data.Result)
	err := rows.Scan(
		&result.ID,
		&result.UserID,
		&result.DiseaseID,
		&result.PredictResult,
		&result.CreatedAt)

	return result, err
}
