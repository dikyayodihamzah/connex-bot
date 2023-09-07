package controller

import (
	"github.com/gofiber/fiber/v2"
)

type MessageController interface {
	NewRouter(app fiber.Router)
}
