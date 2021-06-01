package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": "success"})
	}
}
