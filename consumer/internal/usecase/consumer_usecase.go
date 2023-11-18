package usecase

type consumerUsecase struct {
}

func NewConsumerUsecase() ConsumerUsecase {
	return &consumerUsecase{}
}
