package jwt

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // Tambahkan role ke dalam claim JWT
	jwt.StandardClaims
}
