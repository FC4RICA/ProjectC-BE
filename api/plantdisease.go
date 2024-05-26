package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handlePlantDisease(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreatePlantDiseases(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetPlantDiseases(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreatePlantDiseases(w http.ResponseWriter, r *http.Request) error {
	createPlantDiseasesReq := new(data.CreatePlantDiseasesRequest)
	if err := json.NewDecoder(r.Body).Decode(createPlantDiseasesReq); err != nil {
		return err
	}

	plantDiseases := []data.PlantDisease{}

	for _, createPlantDiseaseReq := range createPlantDiseasesReq.Diseases {
		plantDisease, err := data.NewPlantDisease(createPlantDiseaseReq)
		if err != nil {
			return err
		}

		plantDisease.Plant.ID, plantDisease.Plant.CreatedAt, err = s.store.CreatePlant(plantDisease.Plant)
		if err != nil {
			return err
		}
		plantDisease.Disease.ID, plantDisease.Disease.CreatedAt, err = s.store.CreateDisease(plantDisease.Disease)
		if err != nil {
			return err
		}

		err = s.store.CreatePlantDisease(plantDisease)
		if err != nil {
			return err
		}

		plantDiseases = append(plantDiseases, *plantDisease)
	}

	return util.WriteJSON(w, http.StatusOK, plantDiseases)
}

func (s *APIServer) handleGetPlantDiseases(w http.ResponseWriter, r *http.Request) error {
	plantdiseases, err := s.store.GetPlantDiseases()
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, plantdiseases)
}
