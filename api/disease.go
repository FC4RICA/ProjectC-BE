package api

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleDisease(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "Post" {
		return s.handleCreateDisease(w, r)
	}
	if r.Method == "Get" {
		return s.handleGetDiseases(w, r)
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
