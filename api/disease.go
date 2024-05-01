package api

import (
	"fmt"
	"net/http"
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
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteDisease(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateDisease(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetDiseases(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetDisease(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteDisease(w http.ResponseWriter, r *http.Request) error {
	return nil
}
