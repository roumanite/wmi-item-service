package httpd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/core/domain"
)

type signInPostRequest struct {
	Identifier string
	Password string
}

func (s *Server) SignInPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signInPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err := s.authService.SignIn(domain.SignInRequest{
			Identifier: req.Identifier,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
