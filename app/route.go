package app

import (
	"wmi-item-service/app/handler"
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.POST("/user/sign-up", handler.SignUpPost())

	return router
}