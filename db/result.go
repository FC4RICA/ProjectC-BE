package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateResult(result *data.Result) (int, error) {
	query := `INSERT INTO PredictResult
		(user_id, plant_id, disease_id, created_at)
		VALUES ($1, $2, $3, $4) RETURNING result_id`

	var id int
	err := s.db.QueryRow(
		query,
		result.UserID, result.PlantDisease.Plant.ID, result.PlantDisease.Disease.ID, result.CreatedAt).Scan(&id)
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

		result.PlantDisease, err = s.GetPlantDiseaseByID(result.PlantDisease.Plant.ID, result.PlantDisease.Disease.ID)
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
		result, err := scanIntoResult(rows)
		if err != nil {
			return nil, err
		}

		result.PlantDisease, err = s.GetPlantDiseaseByID(result.PlantDisease.Plant.ID, result.PlantDisease.Disease.ID)
		if err != nil {
			return nil, err
		}

		return result, nil
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

func (s *PostgresStore) DeleteResult(id int) error {
	_, err := s.db.Query("UPDATE PredictResult SET deleted = true, deleted_at = $1 WHERE result_id = $2", time.Now(), id)
	return err
}

func scanIntoResult(rows *sql.Rows) (*data.Result, error) {
	println("error ths here")
	result := new(data.Result)
	result.PlantDisease = new(data.PlantDisease)
	result.PlantDisease.Plant = new(data.Plant)
	result.PlantDisease.Disease = new(data.Disease)
	err := rows.Scan(
		&result.ID,
		&result.UserID,
		&result.PlantDisease.Plant.ID,
		&result.PlantDisease.Disease.ID,
		&result.CreatedAt)
	println("error agin")
	return result, err
}
