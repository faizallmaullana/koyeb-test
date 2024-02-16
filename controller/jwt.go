package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims represents the custom claims for the JWT.
type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

var secretKey = []byte("your-secret-key")

func generateToken(username string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
