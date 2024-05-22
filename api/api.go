package api

import (
	"log"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/db"
	"github.com/Narutchai01/ProjectC-BE/handlers"
	"github.com/Narutchai01/ProjectC-BE/types"
	"github.com/Narutchai01/ProjectC-BE/util"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      db.Storage
}

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			util.WriteJSON(w, http.StatusBadRequest, types.ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/register", makeHTTPHandleFunc(s.handleRegister))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{user-id}", handlers.WithJWTAuth(makeHTTPHandleFunc(s.handleAccountByID), s.store))

	router.HandleFunc("/disease", makeHTTPHandleFunc(s.handleDisease))
	router.HandleFunc("/disease/{user-id}/{disease-id}", handlers.WithJWTAuth(makeHTTPHandleFunc(s.handleDiseaseByID), s.store))

	router.HandleFunc("/result/{user-id}", handlers.WithJWTAuth(makeHTTPHandleFunc(s.handleResult), s.store))
	router.HandleFunc("/result/{user-id}/{result-id}", handlers.WithJWTAuth(makeHTTPHandleFunc(s.handleGetResultByID), s.store))

	log.Println("JSON API server running on port " + s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
