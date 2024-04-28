package handlers

import (
	"fmt"
	"net/http"
)

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")
		handlerFunc(w, r)
	}
}
