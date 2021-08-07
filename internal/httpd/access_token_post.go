package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type accessTokenPostRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (s *Server) AccessTokenPost(expirationMinutes int) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*accessTokenPostRequest)

		claims, err := jwt.ParseToken([]byte(s.jwtKey), req.RefreshToken)

		if err != nil {
			if errors.Is(err, jwt.ExpiredToken) {
				c.Error(ErrExpiredToken)
			} else {
				c.Error(err)
			}
			return
		}

		token, err := jwt.GenerateToken([]byte(s.jwtKey), expirationMinutes, claims.UserId)
		if err != nil {
			c.Error(domain.ErrUnknown)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"accessToken": token,
		})
	}
}
