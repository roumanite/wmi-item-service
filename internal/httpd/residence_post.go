package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type residencePostRequest struct {
	Nickname string `binding:"required"`
	StreetAddress string `json:"streetAddress"`
	City string
	State string
	Country string
	ZipCode string `json:"zipCode"`
	BuildingName string `json:"buildingName"`
}

func (s *Server) ResidencePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req residencePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		residence, err := s.residenceService.CreateResidence(domain.CreateResidenceRequest{
			UserIdOwner: claims.UserId,
			Nickname: req.Nickname,
			StreetAddress: req.StreetAddress,
			City: req.City,
			State: req.State,
			Country: req.Country,
			ZipCode: req.ZipCode,
			BuildingName: req.BuildingName,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success", "residence": residence})
	}
}
