package httpd

import (
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
		metaResults, err := s.residenceService.GetResidenceList(domain.GetResidenceListRequest{
			PerPage: perPage,
			UserIdOwner: claims.UserId,
		})

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, metaResults)
	}
}
