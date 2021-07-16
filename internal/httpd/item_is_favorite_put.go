package httpd

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/httpd/jwt"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type itemIsFavoritePutRequest struct {
	IsFavorite bool `json:"is_favorite"`
}

func (s *Server) ItemIsFavoritePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req itemIsFavoritePutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		err := s.itemService.ToggleIsFavorite(domain.ToggleIsFavoriteRequest{
			RequesterId: claims.UserId,
			ItemId: id,
			IsFavorite: req.IsFavorite,
		})
		if err != nil {
			c.Error(err)
			return
		}

		msg := "Successfully added item to favorite list"
		if !req.IsFavorite {
			msg = "Successfully removed item from favorite list"
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"message": msg,
		})
	}
}