package routes

import (
	"context"

	"go-file-server/db"
	"go-file-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func PublicRoutes(router fiber.Router) {
	router.Get("/", publicIndex)
	router.Get("/metrics", monitor.New(monitor.Config{Title: "file server app moniter"}))
}

func publicIndex(fiberCtx *fiber.Ctx) error {
	response := utils.Response{}
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	_, err := db.Queries.ListUrls(ctx)
	if err != nil {
		response.Msg = "database error"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	return fiberCtx.SendString("public routes")
}
