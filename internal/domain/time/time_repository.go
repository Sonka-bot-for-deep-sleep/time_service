package time

import (
	"context"
	"time"

	"github.com/Sonka-bot-for-deep-sleep/time_service/application/dto"
	"github.com/Sonka-bot-for-deep-sleep/time_service/application/models"
)

type repository interface {
	CreateTime(ctx context.Context, user *dto.CreateUser) error
	GetGoingToBedTime(ctx context.Context, tgID string) (*time.Time, error)
	GetWakeupTime(ctx context.Context, tgID string) (*time.Time, error)
	GetTime(ctx context.Context, tgID string) (*models.Time, error)
	UpdateWakeupTime(ctx context.Context, time dto.UpdateWakeupTime) error
	UpdateGoingToBedTime(ctx context.Context, time dto.UpdateGoingToBedTime) error
}
