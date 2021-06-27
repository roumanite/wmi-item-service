package httpd

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/httpd/jwt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

type myProfilePutRequest struct {
	Bio string
	FirstName string `binding:"required"`
	LastName string
	Birthdate *time.Time
}

func (s *Server) MyProfilePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req myProfilePutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		profile, err := s.userService.UpdateProfile(
			domain.UpdateProfileRequest{
				Id: claims.UserId,
				Bio: req.Bio,
				FirstName: req.FirstName,
				LastName: req.LastName,
				Birthdate: req.Birthdate,
		})

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"profile": profile,
		})
	}
}
