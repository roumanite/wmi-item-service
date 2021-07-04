package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/leebenson/conform"
)

type latestPositionPostRequest struct {
	PositionId int `json:"positionId" binding:"required"`
	LatestPictureUrl string
}

func (s *Server) LatestPositionPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req latestPositionPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		conform.Strings(&req)

		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)
		err := s.itemService.MoveItem(domain.MoveItemRequest{
			UserId: claims.UserId,
			PositionId: req.PositionId,
			LatestPictureUrl: req.LatestPictureUrl,
		})

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"message": "Successfully updated last position of item",
		})
	}
}
