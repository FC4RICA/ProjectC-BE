package api

import (
	"encoding/json"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handleCreateImage(w http.ResponseWriter, r *http.Request) error {
	CreateImageReq := new(data.CreateImageRequest)
	if err := json.NewDecoder(r.Body).Decode(CreateImageReq); err != nil {
		return err
	}

	image, err := data.NewImage(CreateImageReq)
	if err != nil {
		return err
	}
	id, err := s.store.CreateImage(image)
	if err != nil {
		return err
	}
	image.ID = id

	return util.WriteJSON(w, http.StatusOK, image)
}
