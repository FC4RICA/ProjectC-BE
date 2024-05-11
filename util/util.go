package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func Uploadv2(file multipart.File, fileName string) string {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	config := &firebase.Config{
		StorageBucket: "pathfinder-bd7e8.appspot.com",
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf(err.Error())
	}
	storage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	bucket, err := storage.DefaultBucket()
	if err != nil {
		log.Fatalf(err.Error())
	}
	object := bucket.Object("image/" + uuid.NewString() + fileName)
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Fatalf(err.Error())
	}
	if err := wc.Close(); err != nil {
		log.Fatalf(err.Error())
	}
	Name := strings.ReplaceAll(wc.Attrs().Name, "/", "%2F")
	url := "https://firebasestorage.googleapis.com/v0/b/pathfinder-bd7e8.appspot.com/o/" + Name + "?alt=media&token"

	return url
}
