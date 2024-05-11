package data

import (
	"time"
)

type CreateImageRequest struct {
	ResultID    int    `json:"result_id"`
	ImageBase64 string `json:"image_base64"`
}

type Image struct {
	ID        int       `json:"id"`
	ResultID  int       `json:"result_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewImage(image *CreateImageRequest) (*Image, error) {
	//upload img to firebase
	return &Image{
		ResultID:  image.ResultID,
		ImageURL:  "",
		CreatedAt: time.Now().UTC(),
	}, nil
}
