package repository

type consumerRepository struct {
}

func NewConsumerRepository() ConsumerRepository {
	return &consumerRepository{}
}
