package time

import (
	"context"

	"github.com/Sonka-bot-for-deep-sleep/time_service/application/dto"
)

type timeManager struct {
	repo repository
}

func New(repo repository) *timeManager {
	return &timeManager{
		repo: repo,
	}
}

func (t *timeManager) CreateTime(ctx context.Context, req dto.CreateUser) error {
	if err := t.repo.CreateTime(ctx, &req); err != nil {

	}

	return nil
}
