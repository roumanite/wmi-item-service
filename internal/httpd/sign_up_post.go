package httpd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/core/domain"
)

type signUpPostRequest struct {
	Email string `binding:"email,required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Username string `binding:"min=3,max=40,required"`
	Password string `binding:"min=8,required"`
}

func (s *Server) SignUpPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signUpPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err := s.userService.CreateUser(domain.CreateUserRequest{
			Email: req.Email,
			FirstName: req.FirstName,
			LastName: req.LastName,
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
