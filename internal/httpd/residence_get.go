package httpd

// check *****
import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) ResidenceGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		residence, err := s.residenceService.GetResidence(domain.GetResidenceRequest{
			Id: id,
			UserIdOwner: claims.UserId,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success", "residence": residence})
	}
}
