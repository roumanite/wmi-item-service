package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/leebenson/conform"
	"wmi-item-service/internal/core/domain"
)

type signUpPostRequest struct {
	Email string `binding:"email,required" conform:"trim,lower"`
	FirstName string `json:"firstName" binding:"required" conform:"trim"`
	LastName string `json:"lastName" conform:"trim"`
	Username string `binding:"min=3,max=40,required" conform:"trim"`
	Password string `binding:"min=8,required"`
}

func (s *Server) SignUpPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signUpPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		conform.Strings(&req)

		err := s.authService.SignUp(domain.SignUpRequest{
			Email: req.Email,
			FirstName: req.FirstName,
			LastName: req.LastName,
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
