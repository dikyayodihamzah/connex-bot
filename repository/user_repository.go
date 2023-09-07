package repository

import (
	"context"

	"github.com/dikyayodihamzah/connex-bot/model/domain"
)

type UserRepository interface {
	FindAll(c context.Context, param, value string) ([]domain.User, error)
	FindOne(c context.Context, param, value string) (domain.User, error)
	Create(c context.Context, user domain.User) error
	Update(c context.Context, id string, user domain.User) error
	Delete(c context.Context, id string) error
}
