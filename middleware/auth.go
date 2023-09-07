package middleware

import (
	"strings"

	"github.com/dikyayodihamzah/connex-bot/helper"
	"github.com/dikyayodihamzah/connex-bot/model/web"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("token")
	if cookie == "" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	claims, err := helper.ParseJwt(cookie)
	if err != nil {
		if strings.Contains(err.Error(), "Token is expired") {
			return ctx.Status(401).JSON(web.WebResponse{
				Code:    99281,
				Status:  false,
				Message: "Token expired",
			})
		}

		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	ctx.Locals("claims", claims)

	return ctx.Next()
}
