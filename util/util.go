package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/Narutchai01/ProjectC-BE/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func PermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, types.ApiError{Error: "permission denied"})
}

func GetID(r *http.Request, s string) (int, error) {
	idR := s + "-id"
	idStr := mux.Vars(r)[idR]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}

func UploadImageCDN(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	config := &firebase.Config{
		StorageBucket: "pathfinder-bd7e8.appspot.com",
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return "", err
	}
	storage, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}
	bucket, err := storage.DefaultBucket()
	if err != nil {
		return "", err
	}
	object := bucket.Object("image/" + uuid.NewString() + fileName)
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	Name := strings.ReplaceAll(wc.Attrs().Name, "/", "%2F")
	url := "https://firebasestorage.googleapis.com/v0/b/pathfinder-bd7e8.appspot.com/o/" + Name + "?alt=media&token"

	return url, nil
}
