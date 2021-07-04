package httpd

// check *****
import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/leebenson/conform"
)

type residencePutRequest struct {
	Nickname string `binding:"required" conform:"trim"`
	StreetAddress string `json:"streetAddress" conform:"trim"`
	City string `conform:"trim"`
	State string `conform:"trim"`
	Country string `conform:"trim"`
	ZipCode string `json:"zipCode" conform:"trim"`
	BuildingName string `json:"buildingName" conform:"trim"`
}

func (s *Server) ResidencePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req residencePutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		conform.Strings(&req)

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		residence, err := s.residenceService.UpdateResidence(domain.UpdateResidenceRequest{
			Id: id,
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
