package routes

import (
	"go-file-server/db"
	"go-file-server/db/sqlc"
	"go-file-server/middleware"
	"go-file-server/utils"
	"os"
	"strconv"
	"strings"

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

	filename := strings.ReplaceAll(strings.Trim(file.Filename, " "), " ", "")
	fileSize := utils.BytesToMegabytes(file.Size)
	response.Data = filename

	storagePath := os.Getenv("STORAGE")
	storageSize, _ := strconv.Atoi(os.Getenv("STORAGE_SIZE"))

	dicSize, sizeErr := utils.GetFolderSize(storagePath)
	if sizeErr != nil {
		response.Msg = "unable to calculate directory size"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	if dicSize+fileSize > float64(storageSize) {
		response.Msg = "storage size exceeded"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	savingErr := fiberCtx.SaveFile(file, storagePath+filename)
	if savingErr != nil {
		response.Msg = "error during saving file"
		response.StatusCode = fiber.StatusInternalServerError

		return response.CreateJSONResponse(fiberCtx)
	}

	var url sqlc.CreateUrlParams
	if err := fiberCtx.BodyParser(&url); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Queries.CreateUrl(fiberCtx.Context(), url)

	return response.CreateJSONResponse(fiberCtx)
}
