package httpd

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func (s *Server) LivenessGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "success",
			"server_time": time.Now(),
		})
	}
}