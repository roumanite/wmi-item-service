package jwt

import (
	"time"
	"strings"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/lestrrat-go/jwx/jwa"
	"errors"
)

type JwtClaims struct {
	UserId string `json:"userId"`
}

var invalidStructure = errors.New("Invalid JWT structure")
var ExpiredToken = errors.New("JWT has expired")

func ParseToken(jwtKey []byte, tokenString string) (*JwtClaims, error) {
	token, err := jwt.Parse([]byte(tokenString))
	if err != nil {
		return nil, err
	}
	id, ok := token.Get("userId")
	if !ok {
		return nil, invalidStructure
	}
	err = jwt.Validate(token)
	if err != nil {
		if strings.Contains(err.Error(), "exp not satisfied") {
			return nil, ExpiredToken
		}
		return nil, err
	}
	return &JwtClaims{UserId: id.(string)}, nil
}

func GenerateToken(jwtKey []byte, expirationMinutes int, id string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	token := jwt.New()
	token.Set("userId", id)
	token.Set("exp", expirationTime.Unix())
	payload, err := jwt.Sign(token, jwa.HS256, jwtKey)
	if err != nil {
		return "", err
	}
	return string(payload), nil
}
