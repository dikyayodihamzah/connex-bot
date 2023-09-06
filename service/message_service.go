package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dikyayodihamzah/connex-bot/exception"
	"github.com/dikyayodihamzah/connex-bot/model/kafkamodel"
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

func (service *messageServiceImpl) SendMessage(request []byte) error {
	messageRequest := new(kafkamodel.TelegramMessageRequest)

	if err := json.Unmarshal(request, messageRequest); err != nil {
		log.Println("Error kafka send message consumer:", err.Error())
		return err
	}

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
			log.Println(err)
			return exception.NewError(http.StatusInternalServerError, "Failed to send message")
		}
	}

	return nil
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
