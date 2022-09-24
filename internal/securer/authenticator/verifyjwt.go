package authenticator

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("JWT_SECRET_SECRET")

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
				fmt.Println("Invalid JWT token")
			}

			if err != nil {
				fmt.Println("Invalid JWT token")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				client, ok := claims["client"].(string)
				if !ok {
					fmt.Println("Not Authorized: failed to get client from jwt claim")
				}
				if !UserIsValid(client) {
					fmt.Println(fmt.Sprintf("Not Authorized: user %s not found", client))
				}
				next.ServeHTTP(w, r)
			} else {
				fmt.Println("Not Authorized")
			}
		} else {
			fmt.Println("Not Authorized")
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
