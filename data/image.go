package data

import (
	"net/http"
	"time"

	"github.com/Narutchai01/ProjectC-BE/util"
)

type CreateImageRequest struct {
	ResultID int    `json:"result_id"`
	ImageURL string `json:"image_url"`
}

type Image struct {
	ID        int       `json:"id"`
	ResultID  int       `json:"result_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewImage(image *CreateImageRequest) (*Image, error) {
	return &Image{
		ResultID:  image.ResultID,
		ImageURL:  image.ImageURL,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func UploadImages(r *http.Request) ([]string, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return nil, err
	}

	files := r.MultipartForm.File["images"]
	images := []string{}

	for _, file := range files {
		fileHeader, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fileHeader.Close()

		imageURL, err := util.UploadImageCDN(fileHeader, file.Filename)
		if err != nil {
			return nil, err
		}
		images = append(images, imageURL)
	}

	return images, nil
}
