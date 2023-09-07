package model_kafka

import (
	"time"

	"github.com/dikyayodihamzah/connex-bot/model/domain"
)

type User struct {
	Id               string      `json:"id" bson:"_id"`
	Name             string      `json:"name" bson:"name"`
	Username         string      `json:"username" bson:"username"`
	Email            string      `json:"email" bson:"email"`
	Password         string      `json:"password" bson:"password"`
	Phone            string      `json:"phone" bson:"phone"`
	Avatar           string      `json:"avatar" bson:"avatar"`
	RoleId           string      `json:"role_id" bson:"role_id"`
	Role             domain.Role `json:"role" bson:"role"`
	Status           string      `json:"status" bson:"status"`
	FirmwareUser     string      `json:"firmware_user" bson:"firmware_user"`
	FirmwarePassword string      `json:"firmware_password" bson:"firmware_password"`
	CreatedAt        time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time   `json:"update_at" bson:"updated_at"`
	Projects         []string    `json:"projects" bson:"projects"`
	TelegramUser     string      `json:"telegram_user" bson:"telegram_user"`
}
