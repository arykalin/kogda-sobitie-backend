package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arykalin/kogda-sobitie-backend/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(DotEnvVariable("JWT_SECRET"))

// IsAuthorized -> verify jwt header
func IsAuthorized(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if token == nil {
				AuthorizationResponse("Invalid JWT token", w)
			}

			if err != nil {
				AuthorizationResponse("Invalid JWT token", w)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				client, ok := claims["client"].(string)
				if !ok {
					AuthorizationResponse("Not Authorized: failed to get client from jwt claim", w)
				}
				if !auth.UserIsValid(client) {
					AuthorizationResponse(fmt.Sprintf("Not Authorized: user %s not found", client), w)
				}
				next.ServeHTTP(w, r)
			} else {
				AuthorizationResponse("Not Authorized", w)
			}
		} else {
			AuthorizationResponse("Not Authorized", w)
		}
	}
}

// GenerateJWT -> generate jwt
func GenerateJWT(client string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = client
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
