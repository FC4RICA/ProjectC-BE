package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handleResult(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateResult(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetResults(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleResultByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetResultByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteResult(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateResult(w http.ResponseWriter, r *http.Request) error {
	createResultReq := new(data.CreateResultRequest)
	if err := json.NewDecoder(r.Body).Decode(createResultReq); err != nil {
		return err
	}

	result, err := data.NewResult(createResultReq)
	if err != nil {
		return err
	}

	id, err := s.store.CreateResult(result)
	if err != nil {
		return err
	}
	result.ID = id

	return util.WriteJSON(w, http.StatusOK, result)
}

func (s *APIServer) handleGetResultByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteResult(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetResults(w http.ResponseWriter, r *http.Request) error {
	return nil
}
