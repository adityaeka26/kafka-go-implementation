package usecase

import (
	"context"
	"kafka-go-implementation-consumer/internal/dto"
)

type ConsumerUsecase interface {
	TestTopic(ctx context.Context, req dto.TestTopicReq) error
}
