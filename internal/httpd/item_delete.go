package httpd

// check *****
import (
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type itemDeleteRequest struct {
    DeletionNotes string `json:"deletion_notes"`
}

func (s *Server) ItemDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req itemDeleteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, _ := c.Keys[jwtClaimsCtxKey].(JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		err := s.itemService.DeleteItem(domain.DeleteItemRequest{
			Id: id,
			UserIdOwner: claims.UserId,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
