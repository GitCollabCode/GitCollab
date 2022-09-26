package handlers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) (string, error) {
	var sampleSecretKey = []byte("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(48 * time.Hour) // todo update this
	claims["authorized"] = true
	claims["user"] = username
	// create token, return err and token string
	return token.SignedString(sampleSecretKey)
}
