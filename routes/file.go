package routes

import (
	"go-file-server/middleware"
	"go-file-server/utils"

	"github.com/gofiber/fiber/v2"
)

func FileRoutes(router fiber.Router) {
	router.Use(middleware.BasicAuth())

	router.Post("/upload", upload)
}

func upload(fiberCtx *fiber.Ctx) error {
	response := utils.Response{}

	file, err := fiberCtx.FormFile("image")
	if err != nil {
		response.Msg = "no file is selected"
		response.StatusCode = fiber.StatusInternalServerError
		return response.CreateJSONResponse(fiberCtx)
	}

	filepath, err := utils.VeriyStorageAndRetrievePath(file.Filename, file.Size)
	if err != nil {
		response.Msg = "error during saving file"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	response.Data = file.Filename
	savingErr := fiberCtx.SaveFile(file, filepath)
	if savingErr != nil {
		response.Msg = "error during saving file"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	// var url sqlc.CreateUrlParams
	// if err := fiberCtx.BodyParser(&url); err != nil {
	// 	return fiberCtx.Status(fiber.StatusBadRequest).SendString(err.Error())
	// }

	// _, err = db.Queries.CreateUrl(fiberCtx.Context(), url)
	// if err != nil {
	// 	response.Msg = "error during saving to database"
	// 	response.StatusCode = fiber.StatusInternalServerError

	// 	return response.CreateJSONResponse(fiberCtx)
	// }

	return response.CreateJSONResponse(fiberCtx)
}
