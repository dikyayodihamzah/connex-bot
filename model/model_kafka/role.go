package model_kafka

type Role struct {
	Id          string           `json:"_id" bson:"_id"`
	Name        string           `json:"name" bson:"name"`
	Permissions []RolePermission `json:"permissions" bson:"permissions"`
}
