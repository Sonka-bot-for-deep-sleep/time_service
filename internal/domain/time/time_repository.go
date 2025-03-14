package time

import (
	"context"
	"time"

	"github.com/Sonka-bot-for-deep-sleep/time_service/application/dto"
	"github.com/Sonka-bot-for-deep-sleep/time_service/application/models"
)

type repository interface {
	GetGoingToBedTime(ctx context.Context, tgID string) (time.Time, error)
	GetWakeupTime(ctx context.Context, tgID string) (time.Time, error)
	SetWakeupTime(ctx context.Context, time dto.SetWakeupTime) error
	SetGoingToBedTime(ctx context.Context, time dto.SetGoingToBedTime) error
	GetTime(ctx context.Context, tgID string) (*models.Time, error)
}
