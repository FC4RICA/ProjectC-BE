package api

import (
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handleCreateImage(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	files := r.MultipartForm.File["images"]
	imageARR := []string{}

	for _, file := range files {
		fileHeader, err := file.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer fileHeader.Close()

		imageARR = append(imageARR, util.Uploadv2(fileHeader, file.Filename))
	}

	return util.WriteJSON(w, http.StatusOK, imageARR)

}
