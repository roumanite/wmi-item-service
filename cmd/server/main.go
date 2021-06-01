package main

import (
	"wmi-item-service/app"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	server := app.NewServer(router)

	err := server.Run()
	if err != nil {
		return err
	}

	return nil
}
