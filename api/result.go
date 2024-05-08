package api

import (
	"fmt"
	"net/http"
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
	return nil
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
