package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateDisease(dis *data.Disease) (int, time.Time, error) {
	query := `WITH s AS (
			SELECT disease_id
			FROM Disease
			WHERE disease_name = $1
		), i AS (
			INSERT INTO Disease (disease_name, created_at)
			SELECT $1, $2
			WHERE NOT EXISTS (SELECT 1 FROM s)
			RETURNING disease_id
		)
		SELECT disease_id
		FROM i
		UNION ALL
		SELECT disease_id
		FROM s`

	var id int
	var time time.Time
	err := s.db.QueryRow(
		query,
		dis.DiseaseName, dis.CreatedAt).Scan(&id)
	if err != nil {
		return 0, time, err
	}

	return id, time, nil
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

func (s *PostgresStore) GetDiseaseByName(name string) (*data.Disease, error) {
	rows, err := s.db.Query("SELECT * FROM Disease WHERE disease_name = $1", name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoDisease(rows)
	}

	return nil, fmt.Errorf("disease %s not found", name)
}

func scanIntoDisease(rows *sql.Rows) (*data.Disease, error) {
	disease := new(data.Disease)
	err := rows.Scan(
		&disease.ID,
		&disease.DiseaseName,
		&disease.CreatedAt)

	return disease, err
}
