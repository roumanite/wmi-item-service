package httpd

// check *****
import (
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) ItemGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Keys[jwtClaimsCtxKey].(JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		item, err := s.itemService.GetItem(domain.GetItemRequest{
			Id: id,
			UserIdOwner: claims.UserId,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success", "item": item})
	}
}
