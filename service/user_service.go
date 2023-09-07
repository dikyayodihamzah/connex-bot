package service

type UserService interface {
	AddTelegramChatID(chatID int64, username string) error
}
