package service

type UserConsumerService interface {
	Create(message []byte) error
	Update(message []byte) error
	Delete(message []byte) error
}
