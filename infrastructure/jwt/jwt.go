package jwt

import (
	"errors"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type ContextKeyValueType string

const (
	ContextKeyValue ContextKeyValueType = "ContextKeyValue" // TODO: change.
	Key             string              = "Key"             // TODO: change.
)

type Claims struct {
	jwt.StandardClaims
	IsAdmin bool `json:"iad"`
}

func Generate(isAdmin bool) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		IsAdmin: isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
