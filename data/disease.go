package data

import "time"

type Disease struct {
	ID          int         `json:"id"`
	DiseaseName string      `json:"disease_name"`
	PlantName   string      `json:"plant_name"`
	Discription Discription `json:"discription"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Discription struct {
}
