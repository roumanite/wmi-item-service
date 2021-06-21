package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type JwtClaims struct {
	UserId string `json:"userId"`

	jwt.StandardClaims
}

var jwtSigningMethod = jwt.SigningMethodHS256

func ParseToken(jwtKey []byte, tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwtSigningMethod {
			return nil, errors.New("Invalid signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*JwtClaims), nil
}

func GenerateToken(jwtKey []byte, id string) (string, error) {
	expirationTime := time.Now().Add(1000 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &JwtClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
