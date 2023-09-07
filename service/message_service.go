package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dikyayodihamzah/connex-bot/exception"
	"github.com/dikyayodihamzah/connex-bot/model/model_kafka"
	"github.com/dikyayodihamzah/connex-bot/model/web"
	"github.com/dikyayodihamzah/connex-bot/repository"
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageService interface {
	SendMessage(request []byte) error
	TestSendMessage(c context.Context, request web.TelegramMessageRequest) error
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

func (service *messageServiceImpl) TestSendMessage(c context.Context, messageRequest web.TelegramMessageRequest) error {
	validate := validator.New()
	if err := validate.Struct(messageRequest); err != nil {
		return err
	}

	if messageRequest.Message == "" {
		return exception.NewError(http.StatusBadRequest, "Message parameter is missing")
	}

	var userChatIDs []int64
	for _, username := range messageRequest.Usernames {
		ctx := context.Background()

		user, _ := service.UserRepository.FindOne(ctx, "telegram_user", username)
		userChatIDs = append(userChatIDs, user.TelegramChatId)
	}

	for _, chatID := range userChatIDs {
		message := tgbotapi.NewMessage(chatID, messageRequest.Message)
		if _, err := service.Bot.Send(message); err != nil {
			return exception.NewError(http.StatusInternalServerError, "Failed to send message")
		}
	}

	return nil
}

func (service *messageServiceImpl) SendMessage(request []byte) error {
	notificationRequest := new(model_kafka.TelegramNotification)

	if err := json.Unmarshal(request, notificationRequest); err != nil {
		log.Println("Error kafka send message:", err.Error())
		return err
	}

	validate := validator.New()
	if err := validate.Struct(notificationRequest); err != nil {
		return err
	}

	if notificationRequest.Message == "" {
		return exception.NewError(http.StatusBadRequest, "Message parameter is missing")
	}

	user, err := service.UserRepository.FindOne(context.Background(), "telegram_user", notificationRequest.TelegramUser)
	if err != nil {
		return exception.NewError(http.StatusInternalServerError, fmt.Sprintf("Failed to find user with telegram username %s", notificationRequest.TelegramUser))
	}

	message := tgbotapi.NewMessage(user.TelegramChatId, notificationRequest.Message)
	if _, err := service.Bot.Send(message); err != nil {
		return exception.NewError(http.StatusInternalServerError, "Failed to send message")
	}

	return nil
}
