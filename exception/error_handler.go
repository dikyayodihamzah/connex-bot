package exception

import (
	"errors"

	"github.com/dikyayodihamzah/connex-bot/model/web"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	response := web.WebResponse{
		Code:    code,
		Status:  false,
		Message: err.Error(),
	}

	return ctx.Status(code).JSON(response)
}

func NewError(code int, message string) *fiber.Error {
	return fiber.NewError(code, message)
}
