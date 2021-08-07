package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
)

type signInPostRequest struct {
	Identifier string `binding:"required" conform:"trim,lower"`
	Password string `binding:"required"`
}

func (s *Server) SignInPost(atExpirationMinutes int, rtExpirationMinutes int) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*signInPostRequest)

		user, err := s.authService.SignIn(domain.SignInRequest{
			Identifier: req.Identifier,
			Password: req.Password,
		})
		if err != nil {
			c.Error(err)
			return
		}

		atToken, err := jwt.GenerateToken([]byte(s.jwtKey), atExpirationMinutes, user.Id)
		if err != nil {
			c.Error(domain.ErrUnknown)
			return
		}

		rtToken, err := jwt.GenerateToken([]byte(s.jwtKey), rtExpirationMinutes, user.Id)
		if err != nil {
			c.Error(domain.ErrUnknown)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"accessToken": atToken,
			"refreshToken": rtToken,
		})
	}
}
