package httpd

// check *****
import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type itemDeleteRequest struct {
	DeletionNotes string `json:"deletionNotes"`
}

func (s *Server) ItemDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*itemDeleteRequest)

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		err := s.itemService.DeleteItem(domain.DeleteItemRequest{
			Id: id,
			UserIdOwner: claims.UserId,
			DeletionNotes: req.DeletionNotes,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
