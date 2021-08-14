package httpd

import (
	"time"
	"wmi-item-service/internal/httpd/jwt"
	"wmi-item-service/internal/core/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) ResidencesGet() gin.HandlerFunc {
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

		lastId := c.Query("lastId")

		var lastCreatedAt *time.Time
		lca := c.Query("lastCreatedAt")
		if len(lca) > 0 {
			lcaTime, err := time.Parse(time.RFC3339, lca)
			if err == nil {
				lastCreatedAt = &lcaTime
			}
		}

		metaResults, err := s.residenceService.GetResidenceList(domain.GetResidenceListRequest{
			PerPage: perPage,
			UserIdOwner: claims.UserId,
			Order: order,
			LastId: lastId,
			LastCreatedAt: lastCreatedAt,
		})

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, metaResults)
	}
}
