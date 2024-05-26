package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreatePlantDisease(dis *data.PlantDisease) error {
	query := `INSERT INTO PlantDisease 
		(disease_id, plant_id, description, created_at)
		VALUES ($1, $2, $3, $4)`

	bs, err := json.Marshal(dis.Description)
	if err != nil {
		return err
	}

	_, err = s.db.Query(query, dis.Disease.ID, dis.Plant.ID, bs, dis.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetPlantDiseases() ([]*data.PlantDisease, error) {
	rows, err := s.db.Query("SELECT * FROM PlantDisease")
	if err != nil {
		return nil, err
	}

	plantDiseases := []*data.PlantDisease{}
	for rows.Next() {

		plantDisease, err := scanIntoPlantDisease(rows)
		if err != nil {
			return nil, err
		}

		plantDisease.Plant, err = s.GetPlantByID(plantDisease.Plant.ID)
		if err != nil {
			return nil, err
		}
		plantDisease.Disease, err = s.GetDiseaseByID(plantDisease.Disease.ID)
		if err != nil {
			return nil, err
		}

		plantDiseases = append(plantDiseases, plantDisease)
	}

	return plantDiseases, nil
}

func (s *PostgresStore) GetPlantDiseaseByID(plantID, diseaseID int) (*data.PlantDisease, error) {
	rows, err := s.db.Query("SELECT * FROM Disease WHERE plant_id = $1 AND disease_id = $2", plantID, diseaseID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoPlantDisease(rows)
	}

	return nil, fmt.Errorf("disease_id %d, plant_id %d not found", plantID, diseaseID)
}

func scanIntoPlantDisease(rows *sql.Rows) (*data.PlantDisease, error) {
	plantDisease := new(data.PlantDisease)
	plantDisease.Plant = new(data.Plant)
	plantDisease.Disease = new(data.Disease)
	var bs []uint8
	err := rows.Scan(
		&plantDisease.Plant.ID,
		&plantDisease.Disease.ID,
		&bs,
		&plantDisease.CreatedAt)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &plantDisease.Description)

	return plantDisease, err
}
