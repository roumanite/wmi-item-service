package httpd

// check *****
import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type itemPutRequest struct {
    Name string `binding:"required"`
    CategoryId int `json:"categoryId"`
    DisplayPictureUrl string `json:"displayPictureUrl"` 
    Notes string `json:"notes"`
}

func (s *Server) ItemPut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req itemPutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		item, err := s.itemService.UpdateItem(domain.UpdateItemRequest{
			Id: id,
			Name: req.Name,
			UserIdOwner: claims.UserId,
			CategoryId: req.CategoryId,
			DisplayPictureUrl: req.DisplayPictureUrl,
			Notes: req.Notes,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "success", "item": item})
	}
}
