package service

import (
	"context"
	"log"
	"net/http"

	"github.com/dikyayodihamzah/connex-bot/exception"
	"github.com/dikyayodihamzah/connex-bot/model/web"
	"github.com/dikyayodihamzah/connex-bot/repository"
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageService interface {
	Broadcast(c context.Context, request web.MessageRequest) error
}

type messageServiceImpl struct {
	Bot            *tgbotapi.BotAPI
	UserRepository repository.UserRepository
}

func NewMessageService(
	bot *tgbotapi.BotAPI,
	userRepo repository.UserRepository) MessageService {
	return &messageServiceImpl{
		Bot:            bot,
		UserRepository: userRepo,
	}
}

func (service *messageServiceImpl) Broadcast(c context.Context, request web.MessageRequest) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return err
	}

	if request.Message == "" {
		return exception.NewError(http.StatusBadRequest, "Message parameter is missing")
	}

	telegramUsername := "TELEGRAM_BOT_ID"

	user, err := service.UserRepository.FindOne(c, "telegram_user", telegramUsername)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(user.TelegramChatId, request.Message)
	if _, err := service.Bot.Send(message); err != nil {
		log.Println(err)
		return exception.NewError(http.StatusInternalServerError, "Failed to send message")
	}

	return nil
}
