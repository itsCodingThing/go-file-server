package middleware

import (
	"go-file-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "password",
		},
		Unauthorized: func(fiberCtx *fiber.Ctx) error {
			response := utils.Response{
				Msg:        "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
			}

			return response.CreateJSONResponse(fiberCtx)
		},
	})
}
