package httpd

import (
	"reflect"
	"strings"
	"github.com/gin-gonic/gin"
	"wmi-item-service/internal/core/domain"
	"github.com/go-playground/validator/v10"
	"wmi-item-service/internal/translator"
	"github.com/leebenson/conform"
)

func Bind(val interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		panic(`Bind struct can not be a pointer.`)
	}
	typ := value.Type()

	return func(c *gin.Context) {
		obj := reflect.New(typ).Interface()

		if err := c.ShouldBindJSON(obj); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				list := make(map[string]interface{})
				
				for _, fe := range errs {
					list[fe.Field()] = gin.H{
						"type": fe.Tag(),
						"message": translator.Translate(fe, strings.ToLower(c.Request.Header.Get("Accept-Language"))),
					}
				}
				c.Error(domain.CustomError(domain.InvalidRequest, "", list))
			} else {
				c.Error(domain.CustomError(domain.InvalidRequest, err.Error(), nil))
			}
			c.Abort()
			return
		}

		conform.Strings(obj)
		c.Set(gin.BindKey, obj)
	}
}