package httpd

import (
	"wmi-item-service/internal/core/port"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router *gin.Engine
	userService port.UserService
	residenceService port.ResidenceService
}

func NewServer(router *gin.Engine, userService port.UserService, residenceService port.ResidenceService) *Server {
	return &Server{
		router,
		userService,
		residenceService,
	}
}

func (s *Server) Run() error {
	r := s.router

	r.POST("/user/sign-up", s.SignUpPost())
	r.POST("/residence", s.ResidencePost())

	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}