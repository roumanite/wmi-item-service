package httpd

import (
	"wmi-item-service/internal/core/domain"
	"wmi-item-service/internal/httpd/jwt"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type itemIsFavoritePutRequest struct {
	IsFavorite bool `json:"isFavorite"`
}

func (s *Server) ItemIsFavoritePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(gin.BindKey).(*itemIsFavoritePutRequest)

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := s.itemService.ToggleIsFavorite(domain.ToggleIsFavoriteRequest{
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
