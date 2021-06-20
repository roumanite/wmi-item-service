package httpd

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
)

type Err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	jwtNoToken = "jwt-no-token"
	jwtBadToken = "jwt-bad-token"
)

var (
	ErrNoToken = domain.CustomError(jwtNoToken, "Bearer token is not found in HTTP Authorization header", nil)
	ErrBadToken = domain.CustomError(jwtBadToken, "Invalid JWT token or signature", nil)
)

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

		claims, err := jwt.ParseToken(jwtKey, tokenString)

		if err != nil {
			c.Error(ErrBadToken)
			c.Abort()
			return
		}

		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[jwtClaimsCtxKey] = *claims

		c.Next()
	}
}
