package middleware

import (
	model "dizeto-backend/app/model/jwt"
	"fmt"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan token dari header Authorization
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}
		fmt.Println(claims.Role)

		// Periksa role pengguna
		if claims.Role == "admin" {
			// Jika role adalah admin, maka izinkan akses
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Requires admin role"})
		c.Abort()
	}
}
