package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	imagesURL, err := data.UploadImages(r)
	if err != nil {
		return err
	}

	imagesJson, err := json.Marshal(imagesURL)
	if err != nil {
		return err
	}
	body := []byte(fmt.Sprintf(`{
		"imageArr": %s
		}`, string(imagesJson)))

	req, err := http.NewRequest(http.MethodPost, os.Getenv("AI_API"), bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	predictRes := &data.PredictResponse{}
	err = json.NewDecoder(res.Body).Decode(predictRes)
	if err != nil {
		return err
	}

	createResultReq.PlantDisease.Plant, err = s.store.GetPlantByName(predictRes.PlantName)
	if err != nil {
		return err
	}
	createResultReq.PlantDisease.Disease, err = s.store.GetDiseaseByName(predictRes.DiseaseName)
	if err != nil {
		return err
	}
	createResultReq.PlantDisease, err = s.store.GetPlantDiseaseByID(createResultReq.PlantDisease.Plant.ID, createResultReq.PlantDisease.Disease.ID)
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

	for _, imageURL := range imagesURL {
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
