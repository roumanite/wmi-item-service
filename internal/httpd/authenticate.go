package httpd

import (
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"wmi-item-service/internal/core/domain"
)

type JwtClaims struct {
	UserId string `json:"userId"`

	jwt.StandardClaims
}

type Err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	jwtNoToken = "jwt-no-token"
	jwtBadToken = "jwt-bad-token"
)

var (
	ErrNoToken = domain.CustomError(jwtNoToken, "Bearer token is not found in HTTP Authorization header")
	ErrBadToken = domain.CustomError(jwtBadToken, "Invalid JWT token or signature")
)

var jwtSigningMethod = jwt.SigningMethodHS256

const jwtClaimsCtxKey string = "claims"

func (s *Server) Authenticate(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(tokenString, "Bearer ")

		if len(splitToken) <= 1 {
			c.Error(ErrNoToken)
			c.Abort()
			return
		}

		tokenString = splitToken[1]
		token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwtSigningMethod {
				return nil, errors.New("Invalid signing method")
			}
			return jwtKey, nil
		})

		if err != nil {
			c.Error(ErrBadToken)
			c.Abort()
			return
		}

		claims := token.Claims.(*JwtClaims)
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[jwtClaimsCtxKey] = *claims

		c.Next()
	}
}
