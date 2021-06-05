package httpd

import (
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type residencePostRequest struct {
	Nickname string
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

		s.residenceService.CreateResidence(domain.CreateResidenceRequest{
			Nickname: req.Nickname,
			StreetAddress: req.StreetAddress,
			City: req.City,
			State: req.State,
			Country: req.Country,
			ZipCode: req.ZipCode,
			BuildingName: req.BuildingName,
		})
		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
