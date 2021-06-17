package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/core/domain"
)

var codeToStatus = map[string]int{
	domain.NotFound: http.StatusNotFound,
	domain.Unknown: http.StatusInternalServerError,
	jwtNoToken: http.StatusUnauthorized,
	jwtBadToken: http.StatusUnauthorized,
}

func handleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

    errorToReturn := c.Errors.Last()
    if errorToReturn != nil {
			if err, ok := errorToReturn.Err.(domain.CustomErr); ok {
				status := codeToStatus[err.Code()]
				c.JSON(status, gin.H{
					"code": err.Code(),
					"message": err.Error(),
				})
			}
    }
	}
}
