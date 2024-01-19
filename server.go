package main

import (
	"context"
	"os"

	"go-file-server/db"
	"go-file-server/routes"
	"go-file-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	utils.LoadEnv()

	conn := db.Connect()
	defer conn.Close(context.Background())

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	apiV1 := app.Group("/api/v1")

	routes.PublicRoutes(apiV1)
	routes.FileRoutes(apiV1)

	port, isExist := os.LookupEnv("PORT")
	if !isExist {
		port = "3001"
	}

	app.Listen(":" + port)
}
