package httpd

import (
	"wmi-item-service/internal/core/port"
	"wmi-item-service/internal/translator"
	"log"

	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

func (s *Server) Run(expirationMinutes int) error {
	r := s.router

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		translator.RegisterTranslations(v)
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" || len(name) == 0 {
				return fld.Name
			}

			return name
		})
	}

	r.Use(respondWithError())
	r.POST("/user/sign-up", Bind(signUpPostRequest{}), s.SignUpPost())
	r.POST("/user/sign-in", s.SignInPost(expirationMinutes))

	loginRequired := r.Group(".")
	loginRequired.Use(s.Authenticate([]byte(s.jwtKey)))
	{
		loginRequired.POST("/residence", s.ResidencePost())
		loginRequired.POST("/item", Bind(itemPostRequest{}), s.ItemPost())
		loginRequired.POST("/item/:id/latest-position", s.LatestPositionPost())

		loginRequired.PUT("/user/my-profile", s.MyProfilePut())
		loginRequired.PUT("/residence/:id", s.ResidencePut())
		loginRequired.PUT("/item/:id", s.ItemPut())
		loginRequired.PUT("/item/:id/is-favorite", s.ItemIsFavoritePut())

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
