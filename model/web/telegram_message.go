package web

type TelegramMessageRequest struct {
	UserID  []int64 `json:"user_id"`
	Message string  `json:"message"`
}
