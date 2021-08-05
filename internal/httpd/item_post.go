package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type itemPostRequest struct {
    Name string `json:"name" binding:"required" conform:"trim"`
    CategoryId int `json:"categoryId" conform:"trim"`
    DisplayPictureUrl string `json:"displayPictureUrl" conform:"trim"` 
    Notes string `json:"notes"`
}

func (s *Server) ItemPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*itemPostRequest)

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		item, err := s.itemService.CreateItem(domain.CreateItemRequest{
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
