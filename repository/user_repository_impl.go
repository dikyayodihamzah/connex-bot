package repository

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dikyayodihamzah/connex-bot/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collUser   = os.Getenv("USER_COLLECTION")
	timeout, _ = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
)

type userRepositoryImpl struct {
	CollUser *mongo.Collection
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		CollUser: database.Collection(collUser),
	}
}

func (repository *userRepositoryImpl) FindAll(c context.Context, param, value string) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(timeout)*time.Second)
	defer cancel()

	var users []domain.User
	filter := bson.M{param: value}
	cursor, err := repository.CollUser.Find(ctx, filter)
	if err != nil {
		log.Println("Error while finding users: ", err)
		return users, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		log.Println("Error while decoding users: ", err)
		return users, err
	}

	log.Println("Success found all users")
	return users, nil
}

func (repository *userRepositoryImpl) FindOne(c context.Context, param, value string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(timeout)*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.M{param: value}
	if err := repository.CollUser.FindOne(ctx, filter).Decode(&user); err != nil {
		log.Println("Error while finding user: ", err)
		return user, err
	}

	log.Println("Success found one user")
	return user, nil
}

func (repository *userRepositoryImpl) Create(c context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(timeout)*time.Second)
	defer cancel()

	if _, err := repository.CollUser.InsertOne(ctx, user); err != nil {
		log.Println("Error while creating user: ", err)
		return err
	}

	log.Println("Success created user")
	return nil
}

func (repository *userRepositoryImpl) Update(c context.Context, id string, user domain.User) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(timeout)*time.Second)
	defer cancel()

	update := bson.M{
		"name":              user.Name,
		"username":          user.Username,
		"email":             user.Email,
		"password":          user.Password,
		"phone":             user.Phone,
		"avatar":            user.Avatar,
		"role_id":           user.RoleId,
		"role":              user.Role,
		"status":            user.Status,
		"firmware_user":     user.FirmwareUser,
		"firmware_password": user.FirmwarePassword,
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
		"projects":          user.Projects,
		"telegram_user":     user.TelegramUser,
		"telegram_chat_id":  user.TelegramChatId,
	}

	if _, err := repository.CollUser.UpdateByID(ctx, id, bson.M{"$set": update}); err != nil {
		log.Println("Error while updating user: ", err)
		return err
	}

	log.Println("Success updated user")
	return nil
}

func (repository *userRepositoryImpl) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, time.Duration(timeout)*time.Second)
	defer cancel()

	if _, err := repository.CollUser.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		log.Println("Error while deleting user: ", err)
		return err
	}

	log.Println("Success deleted user")
	return nil
}
