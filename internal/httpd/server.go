package httpd

import (
	"wmi-item-service/internal/core/port"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router *gin.Engine
	authService port.AuthService
	userService port.UserService
	residenceService port.ResidenceService
	itemService port.ItemService
	jwtKey string
}

func NewServer(
	router *gin.Engine,
	authService port.AuthService,
	userService port.UserService,
	residenceService port.ResidenceService,
	itemService port.ItemService,
	jwtKey string,
) *Server {
	return &Server{
		router,
		authService,
		userService,
		residenceService,
		itemService,
		jwtKey,
	}
}

func (s *Server) Run() error {
	r := s.router

	r.Use(handleError())
	r.POST("/user/sign-up", s.SignUpPost())
	r.POST("/user/sign-in", s.SignInPost())

	loginRequired := r.Group(".")
	loginRequired.Use(s.Authenticate([]byte(s.jwtKey)))
	{
		loginRequired.POST("/residence", s.ResidencePost())
		loginRequired.POST("/item", s.ItemPost())

		loginRequired.PUT("/residence/:id", s.ResidencePut())
		loginRequired.PUT("/item/:id", s.ItemPut())

		loginRequired.GET("/user/my-profile", s.MyProfileGet())
		loginRequired.GET("/residence", s.ResidencesGet())
		loginRequired.GET("/residence/:id", s.ResidenceGet())
		loginRequired.GET("/item", s.ItemsGet())
		loginRequired.GET("/item/:id", s.ItemGet())

		loginRequired.DELETE("/residence/:id", s.ResidenceDelete())
		loginRequired.DELETE("/item/:id", s.ItemDelete())
	}

	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
