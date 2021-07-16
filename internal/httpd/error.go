package httpd

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/core/domain"
)

var codeToStatus = map[string]int{
	domain.NotFound: http.StatusNotFound,
	domain.Unknown: http.StatusInternalServerError,
	domain.InvalidRequest: http.StatusBadRequest,
	jwtNoToken: http.StatusUnauthorized,
	jwtBadToken: http.StatusUnauthorized,
	jwtExpired: http.StatusUnauthorized,
}

func respondWithError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

    errorToReturn := c.Errors.Last()
    if errorToReturn != nil {
			if err, ok := errorToReturn.Err.(*domain.CustomErr); ok {
				status := codeToStatus[err.Code()]
				if len(err.Details()) > 0 {
					c.JSON(status, gin.H{
						"code": err.Code(),
						"message": err.Error(),
						"details": err.Details(),
					})
				} else {
					c.JSON(status, gin.H{
						"code": err.Code(),
						"message": err.Error(),
					})
				}
			} else {
				c.JSON(500, gin.H{
					"code": "unknown",
					"message": errorToReturn.Error(),
				})
			}
    }
	}
}
