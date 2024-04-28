package handlers

import (
	"fmt"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/types"
	"github.com/Narutchai01/ProjectC-BE/util"
	jwt "github.com/golang-jwt/jwt/v5"
)

const jwtkey = "testtest"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		_, err := validateJWT(tokenString)
		if err != nil {
			util.WriteJSON(w, http.StatusForbidden, types.ApiError{Error: "Invalid token"})
			return
		}

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtkey), nil
	})
}
