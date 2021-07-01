package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
)

type signInPostRequest struct {
	Identifier string
	Password string
}

func (s *Server) SignInPost(expirationMinutes int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signInPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user, err := s.authService.SignIn(domain.SignInRequest{
			Identifier: req.Identifier,
			Password: req.Password,
		})
		if err != nil {
			c.Error(err)
			return
		}

		token, err := jwt.GenerateToken([]byte(s.jwtKey), expirationMinutes, user.Id)
		if err != nil {
			c.Error(domain.ErrUnknown)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"token": token,
		})
	}
}
