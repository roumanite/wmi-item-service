package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}