package service

import (
	"context"

	"github.com/dikyayodihamzah/connex-bot/model/web"
)

type MessageService interface {
	SendMessage(request []byte) error
	TestSendMessage(c context.Context, request web.TelegramMessageRequest) error
}
