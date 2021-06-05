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
	itemService port.ItemService
	jwtKey string
}

func NewServer(router *gin.Engine, userService port.UserService, residenceService port.ResidenceService, itemService port.ItemService, jwtKey string) *Server {
	return &Server{
		router,
		userService,
		residenceService,
		itemService,
		jwtKey,
	}
}

func (s *Server) Run() error {
	r := s.router

	r.POST("/user/sign-up", s.SignUpPost())

	loginRequired := r.Group(".")
	loginRequired.Use(s.Authenticate([]byte(s.jwtKey)))
	{
		loginRequired.POST("/residence", s.ResidencePost())
		loginRequired.POST("/item", s.ItemPost())
	}

	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
