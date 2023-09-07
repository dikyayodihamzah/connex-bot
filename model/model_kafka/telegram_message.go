package model_kafka

type TelegramMessageRequest struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type TelegramNotification struct {
	UserId       string `query:"user_id" validate:"required"`
	ProjectId    string `query:"project_id"`
	GroupId      string `query:"group_id"`
	NodeId       string `query:"node_id" validate:"required"`
	SensorId     string `query:"sensor_id"`
	ActuatorId   string `query:"actuator_id"`
	Phone        string `query:"phone"`
	Email        string `query:"email"`
	Subject      string `query:"subject"`
	Message      string `query:"message" validate:"required"`
	Level        string `query:"level" validate:"required"`
	Type         string `query:"type"`
	TelegramUser string `query:"telegram_user"`
}
