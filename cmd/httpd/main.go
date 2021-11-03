package main

import (
	"wmi-item-service/internal/config"
	"wmi-item-service/internal/httpd"
	"wmi-item-service/internal/repository"
	"wmi-item-service/internal/core/service"
	"wmi-item-service/pkg/postgres"
	"wmi-item-service/internal/translator"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/database/postgres"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %s\n", err)
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

	if err != nil {
		return err
	}

	err = translator.LoadTranslations("translations/common", "translations/validation")
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepo(dbConn)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	residenceRepo := repository.NewResidenceRepo(dbConn)
	residenceService := service.NewResidenceService(residenceRepo)

	itemRepo := repository.NewItemRepo(dbConn)
	itemService := service.NewItemService(itemRepo)

	server := httpd.NewServer(router, authService, userService, residenceService, itemService, cfg.JwtKey)

	err = server.Run(cfg.AccessTokenExpirationInMinutes, cfg.RefreshTokenExpirationInMinutes)
	if err != nil {
		return err
	}

	return nil
}
