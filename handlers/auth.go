package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/db"
	"github.com/Narutchai01/ProjectC-BE/util"
	jwt "github.com/golang-jwt/jwt/v5"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			util.PermissionDenied(w)
			return
		}
		if !token.Valid {
			util.PermissionDenied(w)
			return
		}
		userID, err := util.GetID(r)
		if err != nil {
			util.PermissionDenied(w)
			return
		}
		account, err := s.GetAccountByID(userID)
		if err != nil {
			util.PermissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		if account.ID != int(claims["AccountID"].(float64)) {
			util.PermissionDenied(w)
			return
		}

		handlerFunc(w, r)

	}
}

func CreateJWT(account *data.Account) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Unix(1516239022, 0)),
		"AccountID": account.ID,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
