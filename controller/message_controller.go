package controller

import (
	"github.com/dikyayodihamzah/connex-bot/exception"
	"github.com/dikyayodihamzah/connex-bot/middleware"
	"github.com/dikyayodihamzah/connex-bot/model/web"
	"github.com/dikyayodihamzah/connex-bot/service"
	"github.com/gofiber/fiber/v2"
)

type MessageController interface {
	NewRouter(app fiber.Router)
}

type messageControllerImpl struct {
	MessageService service.MessageService
}

func NewMessageController(
	msgService service.MessageService) MessageController {
	return &messageControllerImpl{
		MessageService: msgService,
	}
}

func (controller *messageControllerImpl) NewRouter(app fiber.Router) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	app.Post("/messages", middleware.IsAuthenticated, controller.SendMessage)
}

func (controller *messageControllerImpl) SendMessage(ctx *fiber.Ctx) error {
	request := new(web.TelegramMessageRequest)
	if err := ctx.BodyParser(request); err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	if err := controller.MessageService.TestSendMessage(ctx.Context(), *request); err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Message sent successfully",
	})
}
