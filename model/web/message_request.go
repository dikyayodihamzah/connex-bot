package web

type MessageRequest struct {
	Message string `json:"message" form:"message"`
}
