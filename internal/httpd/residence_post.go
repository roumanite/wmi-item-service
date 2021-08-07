package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type residencePostRequest struct {
	Nickname string `binding:"required" conform:"trim"`
	StreetAddress string `json:"streetAddress" conform:"trim"`
	City string `conform:"trim"`
	State string `conform:"trim"`
	Country string `conform:"trim"`
	ZipCode string `json:"zipCode" conform:"trim"`
	BuildingName string `json:"buildingName" conform:"trim"`
}

func (s *Server) ResidencePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*residencePostRequest)

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
