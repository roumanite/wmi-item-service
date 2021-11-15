package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *Server) LivenessGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "success",
		})
	}
}