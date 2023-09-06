package main

import (
	"log"

	"github.com/dikyayodihamzah/connex-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)

// The chat ID of the group or channel where you want to broadcast the message
const chatID = "@your_channel_or_group"

func main() {
	app := fiber.New()

	bot := config.NewBot()

	app.Post("/broadcast", func(c *fiber.Ctx) error {
		messageText := c.Query("message")
		if messageText == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Message parameter is missing")
		}

		msg := tgbotapi.NewMessageToChannel(chatID, messageText)
		_, err := bot.Send(msg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to send message")
		}

		return c.SendString("Message sent successfully")
	})

	log.Println("Server is running on :8080")
	log.Fatal(app.Listen(":8080"))
}
