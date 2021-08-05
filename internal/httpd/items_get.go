package httpd

import (
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) ItemsGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Keys[jwtClaimsCtxKey].(jwt.JwtClaims)

		perPage, _ := strconv.Atoi(c.Query("perPage"))
		if perPage <= 0 {
			perPage = 200
		}

		order := c.DefaultQuery("order", "desc")
		if order != "asc" && order != "desc" {
			order = "desc"
		}

		metaResults, err := s.itemService.GetItemList(domain.GetItemListRequest{
			PerPage: perPage,
			Order: order,
			UserIdOwner: claims.UserId,
		})

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, metaResults)
	}
}
