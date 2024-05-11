package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateDisease(dis *data.Disease) (int, error) {
	query := `INSERT INTO Disease 
		(disease_name, plant_name, description, created_at)
		VALUES ($1, $2, $3, $4) RETURNING  disease_id`

	bs, err := json.Marshal(dis.Description)
	if err != nil {
		return -1, nil
	}

	var id int
	err = s.db.QueryRow(
		query,
		dis.DiseaseName, dis.PlantName, bs, dis.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PostgresStore) GetDiseases() ([]*data.Disease, error) {
	rows, err := s.db.Query("SELECT * FROM Disease")
	if err != nil {
		return nil, err
	}

	diseases := []*data.Disease{}
	for rows.Next() {
		disease, err := scanIntoDisease(rows)
		if err != nil {
			return nil, err
		}

		diseases = append(diseases, disease)
	}

	return diseases, nil
}

func (s *PostgresStore) GetDiseaseByID(id int) (*data.Disease, error) {
	rows, err := s.db.Query("SELECT * FROM Disease WHERE disease_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoDisease(rows)
	}

	return nil, fmt.Errorf("disease %d not found", id)
}

func (s *PostgresStore) GetDiseaseByName(plantName, diseaseName string) (*data.Disease, error) {
	rows, err := s.db.Query("SELECT * FROM Disease WHERE plant_name = $1, disease_name = $2", plantName, diseaseName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoDisease(rows)
	}

	return nil, fmt.Errorf("disease %s %s not found", plantName, diseaseName)
}

func scanIntoDisease(rows *sql.Rows) (*data.Disease, error) {
	disease := new(data.Disease)
	var bs []uint8
	err := rows.Scan(
		&disease.ID,
		&disease.DiseaseName,
		&disease.PlantName,
		&bs,
		&disease.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &disease.Description)

	return disease, err
}
