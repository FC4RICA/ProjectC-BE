package db

import (
	"database/sql"
	"fmt"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateImage(img *data.Image) (int, error) {
	query := `INSERT INTO Image 
		(result_id, image_url, created_at)
		VALUES ($1, $2, $3) RETURNING  image_id`

	var id int
	err := s.db.QueryRow(
		query,
		img.ResultID, img.ImageURL, img.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PostgresStore) GetImages() ([]*data.Image, error) {
	rows, err := s.db.Query("SELECT * FROM Image")
	if err != nil {
		return nil, err
	}

	images := []*data.Image{}
	for rows.Next() {
		image, err := scanIntoImage(rows)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (s *PostgresStore) GetImagesByResultID(id int) ([]*data.Image, error) {
	rows, err := s.db.Query("SELECT * FROM Image WHERE result_id = $1", id)
	if err != nil {
		return nil, err
	}

	images := []*data.Image{}
	for rows.Next() {
		image, err := scanIntoImage(rows)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (s *PostgresStore) GetImageByID(id int) (*data.Image, error) {
	rows, err := s.db.Query("SELECT * FROM Image WHERE image_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoImage(rows)
	}

	return nil, fmt.Errorf("image %d not found", id)
}

func (s *PostgresStore) DeleteImage(id int) error {
	_, err := s.db.Query("DELETE FROM Image WHERE image_Id = $1", id)
	return err
}

func scanIntoImage(rows *sql.Rows) (*data.Image, error) {
	image := new(data.Image)
	err := rows.Scan(
		&image.ID,
		&image.ResultID,
		&image.ImageURL,
		&image.CreatedAt)

	return image, err
}
