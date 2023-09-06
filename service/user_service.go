package service

import (
	"context"

	"github.com/dikyayodihamzah/connex-bot/repository"
)

type UserService interface {
	AddTelegramChatID(chatID int64, username string) error
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepo,
	}
}

func (service *userServiceImpl) AddTelegramChatID(chatID int64, username string) error {
	ctx := context.Background()

	user, err := service.UserRepository.FindOne(ctx, "telegram_user", username)
	if err != nil {
		return err
	}

	user.TelegramChatId = chatID

	if err := service.UserRepository.Update(ctx, user.Id, user); err != nil {
		return err
	}

	return nil
}
