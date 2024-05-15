package api

import (
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
		return s.handleGetResultsByUserID(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateResult(w http.ResponseWriter, r *http.Request) error {
	createResultReq := new(data.CreateResultRequest)
	userid, err := util.GetID(r, "user")
	if err != nil {
		return err
	}
	createResultReq.UserID = userid

	createResultReq.Images, err = data.UploadImages(r)
	if err != nil {
		return err
	}

	result, err := data.NewResult(createResultReq)
	if err != nil {
		return err
	}

	result.ID, err = s.store.CreateResult(result)
	if err != nil {
		return err
	}

	for _, imageURL := range createResultReq.Images {
		createImageReq := &data.CreateImageRequest{
			ResultID: result.ID,
			ImageURL: imageURL,
		}
		image, err := data.NewImage(createImageReq)
		if err != nil {
			return err
		}

		_, err = s.store.CreateImage(image)
		if err != nil {
			return err
		}
	}

	return util.WriteJSON(w, http.StatusOK, result)
}

func (s *APIServer) handleGetResultsByUserID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	id, err := util.GetID(r, "user")
	if err != nil {
		return err
	}

	results, err := s.store.GetResultsByUserID(id)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, results)
}

func (s *APIServer) handleGetResultByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	id, err := util.GetID(r, "result")
	if err != nil {
		return err
	}

	result, err := s.store.GetResultByID(id)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, result)
}
