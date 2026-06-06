package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Parse(tokenString string, secret string) (isValid bool, token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return false, nil, fmt.Errorf("%s: %w", "error parsing token", err)
	}

	if token.Valid {
		return true, token, nil
	}
	return false, nil, nil
}

func GetUserClaims(token *jwt.Token) (UserClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id, ok := claims["uuid"].(string)
		if !ok {
			return UserClaims{}, fmt.Errorf("error getting user id")
		}
		email, ok := claims["email"].(string)
		if !ok {
			return UserClaims{}, fmt.Errorf("error getting user email")
		}

		return UserClaims{
			ID:    id,
			Email: email,
		}, nil
	}
	return UserClaims{}, fmt.Errorf("error getting user claims")
}
