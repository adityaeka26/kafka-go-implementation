package usecase

import (
	"context"
	"kafka-go-implementation-consumer/internal/dto"
	"kafka-go-implementation-consumer/internal/repository"
	"log"
)

type consumerUsecase struct {
	consumerRepository repository.ConsumerRepository
}

func NewConsumerUsecase(consumerRepository repository.ConsumerRepository) ConsumerUsecase {
	return &consumerUsecase{
		consumerRepository: consumerRepository,
	}
}

func (u *consumerUsecase) TestTopic(ctx context.Context, req dto.TestTopicReq) error {
	log.Println(req)

	return nil
}
