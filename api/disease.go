package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handleDisease(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateDisease(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetDiseases(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDiseaseByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetDiseaseByID(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateDisease(w http.ResponseWriter, r *http.Request) error {
	createDiseaseReq := new(data.CreateDiseaseRequest)
	if err := json.NewDecoder(r.Body).Decode(createDiseaseReq); err != nil {
		return err
	}

	disease, err := data.NewDisease(createDiseaseReq)
	if err != nil {
		return err
	}

	disease.ID, err = s.store.CreateDisease(disease)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, disease)
}

func (s *APIServer) handleGetDiseases(w http.ResponseWriter, r *http.Request) error {
	diseases, err := s.store.GetDiseases()
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, diseases)
}

func (s *APIServer) handleGetDiseaseByID(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetID(r, "disease")
	if err != nil {
		return err
	}

	disease, err := s.store.GetDiseaseByID(id)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, disease)
}
