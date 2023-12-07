package routes

import (
	"context"
	"fmt"
	"go-file-server/db"
	"go-file-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(router fiber.Router) {
	router.Get("/", publicIndex)
}

func publicIndex(fiberCtx *fiber.Ctx) error {
	response := utils.Response{}
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	urls, err := db.Queries.ListUrls(ctx)
	if err != nil {
		response.Msg = "no file is selected"
		response.StatusCode = fiber.StatusInternalServerError
		return response.CreateJSONResponse(fiberCtx)
	}

	for _, url := range urls {
		fmt.Println(url)
	}

	return fiberCtx.SendString("public routes")
}
