package httpd

import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type JwtClaims struct {
	UserId string `json:"userId"`

	jwt.StandardClaims
}

type Err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var jwtSigningMethod = jwt.SigningMethodHS256

const jwtClaimsCtxKey string = "claims"

func (s *Server) Authenticate(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(tokenString, "Bearer ")

		if len(splitToken) <= 1 {
			c.JSON(http.StatusUnauthorized, Err{
					Code:    "jwt-no-token",
					Message: "Bearer token is not found in HTTP Authorization header. ",
			})

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
			c.JSON(http.StatusUnauthorized, Err{
				Code:    "jwt-bad-token",
				Message: "Invalid JWT token or signature.",
			})

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
