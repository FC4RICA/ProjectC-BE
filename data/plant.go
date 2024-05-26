package data

import "time"

type Plant struct {
	ID        int       `json:"id"`
	PlantName string    `json:"plant_name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePlantRequest struct {
	PlantName string `json:"plant_name"`
}

func NewPlant(plant *CreatePlantRequest) (*Plant, error) {
	return &Plant{
		PlantName: plant.PlantName,
		CreatedAt: time.Now().UTC(),
	}, nil
}
