package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreatePlant(dis *data.Plant) (int, time.Time, error) {
	query := `WITH s AS (
			SELECT plant_id, created_at
			FROM Plant
			WHERE plant_name = $1
		), i AS (
			INSERT INTO Plant (plant_name, created_at)
			SELECT $1, $2
			WHERE NOT EXISTS (SELECT 1 FROM s)
			RETURNING plant_id, created_at
		)
		SELECT plant_id, created_at
		FROM i
		UNION ALL
		SELECT plant_id, created_at
		FROM s`

	var id int
	var time time.Time
	err := s.db.QueryRow(
		query,
		dis.PlantName, dis.CreatedAt).Scan(&id, &time)
	if err != nil {
		return 0, time, err
	}

	return id, time, nil
}

func (s *PostgresStore) GetPlantByID(id int) (*data.Plant, error) {
	rows, err := s.db.Query("SELECT * FROM Plant WHERE plant_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoPlant(rows)
	}

	return nil, fmt.Errorf("plant %d not found", id)
}

func (s *PostgresStore) GetPlantByName(name string) (*data.Plant, error) {
	rows, err := s.db.Query("SELECT * FROM Plant WHERE plant_name = $1", name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoPlant(rows)
	}

	return nil, fmt.Errorf("plant %s not found", name)
}

func scanIntoPlant(rows *sql.Rows) (*data.Plant, error) {
	plant := new(data.Plant)
	err := rows.Scan(
		&plant.ID,
		&plant.PlantName,
		&plant.CreatedAt)

	return plant, err
}
