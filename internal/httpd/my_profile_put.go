package httpd

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/httpd/jwt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

type myProfilePutRequest struct {
	Bio string `conform:"trim"`
	FirstName string `binding:"required" conform:"trim"`
	LastName string `conform:"trim"`
	Birthdate *time.Time
}

func (s *Server) MyProfilePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*myProfilePutRequest)

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
