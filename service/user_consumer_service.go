package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dikyayodihamzah/connex-bot/model/domain"
	"github.com/dikyayodihamzah/connex-bot/model/model_kafka"
	"github.com/dikyayodihamzah/connex-bot/repository"
)

type UserConsumerService interface {
	Create(message []byte) error
	Update(message []byte) error
	Delete(message []byte) error
}

type userConsumerServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserConsumerService(userRepo repository.UserRepository) UserConsumerService {
	return &userConsumerServiceImpl{
		UserRepository: userRepo,
	}
}

func (service *userConsumerServiceImpl) Create(message []byte) error {
	userMsg := new(model_kafka.User)

	if err := json.Unmarshal(message, userMsg); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	user := domain.User{
		Id:               userMsg.Id,
		Name:             userMsg.Name,
		Username:         userMsg.Username,
		Email:            userMsg.Email,
		Password:         userMsg.Password,
		Phone:            userMsg.Phone,
		Avatar:           userMsg.Avatar,
		RoleId:           userMsg.RoleId,
		Role:             userMsg.Role,
		Status:           userMsg.Status,
		FirmwareUser:     userMsg.FirmwareUser,
		FirmwarePassword: userMsg.FirmwarePassword,
		CreatedAt:        userMsg.CreatedAt,
		UpdatedAt:        userMsg.UpdatedAt,
		Projects:         userMsg.Projects,
		TelegramUser:     userMsg.TelegramUser,
	}

	ctx := context.Background()
	if err := service.UserRepository.Create(ctx, user); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	return nil
}

func (service *userConsumerServiceImpl) Update(message []byte) error {
	userMsg := new(model_kafka.User)

	if err := json.Unmarshal(message, userMsg); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	user := domain.User{
		Id:               userMsg.Id,
		Name:             userMsg.Name,
		Username:         userMsg.Username,
		Email:            userMsg.Email,
		Password:         userMsg.Password,
		Phone:            userMsg.Phone,
		Avatar:           userMsg.Avatar,
		RoleId:           userMsg.RoleId,
		Role:             userMsg.Role,
		Status:           userMsg.Status,
		FirmwareUser:     userMsg.FirmwareUser,
		FirmwarePassword: userMsg.FirmwarePassword,
		CreatedAt:        userMsg.CreatedAt,
		UpdatedAt:        userMsg.UpdatedAt,
		Projects:         userMsg.Projects,
		TelegramUser:     userMsg.TelegramUser,
	}

	ctx := context.Background()
	if err := service.UserRepository.Update(ctx, user.Id, user); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	return nil
}

func (service *userConsumerServiceImpl) Delete(message []byte) error {
	userMsg := new(model_kafka.User)

	if err := json.Unmarshal(message, userMsg); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	ctx := context.Background()
	if err := service.UserRepository.Delete(ctx, userMsg.Id); err != nil {
		log.Println("Error kafka insert user consumer:", err.Error())
		return err
	}

	return nil
}
