package main

import (
	"wmi-item-service/internal/config"
	"wmi-item-service/internal/httpd"
	"wmi-item-service/internal/repository"
	"wmi-item-service/internal/core/service"
	"wmi-item-service/pkg/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/database/postgres"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.LoadConfig()

	router := gin.Default()

	dbConn, err := postgres.NewConnection(
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Port,
		cfg.Database.Schema,
	)

	userRepo := repository.NewUserRepo(dbConn)
	authService := service.NewAuthService(userRepo)

	residenceRepo := repository.NewResidenceRepo(dbConn)
	residenceService := service.NewResidenceService(residenceRepo)

	itemRepo := repository.NewItemRepo(dbConn)
	itemService := service.NewItemService(itemRepo)

	server := httpd.NewServer(router, authService, residenceService, itemService, cfg.JwtKey)

	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}


