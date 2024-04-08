package utils

import (
	model "dizeto-backend/app/model/jwt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

func GenerateJWT(username, password, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 1 hari

	claims := &model.Claims{
		Username: username,
		Password: password,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
