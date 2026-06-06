package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID    string
	Email string
}

func Sign(userClaims UserClaims, ttl time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = userClaims.ID
	claims["email"] = userClaims.Email
	claims["exp"] = time.Now().Add(ttl).Unix()

	return token.SignedString([]byte(secret))
}
