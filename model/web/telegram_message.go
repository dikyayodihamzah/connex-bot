package web

type TelegramMessageRequest struct {
	Usernames []string `json:"usernames"`
	Message   string   `json:"message"`
}
