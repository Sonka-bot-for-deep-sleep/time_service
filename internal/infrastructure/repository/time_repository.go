package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Sonka-bot-for-deep-sleep/time_service/application/dto"
	"github.com/Sonka-bot-for-deep-sleep/time_service/application/models"
	"github.com/jackc/pgx/v5"
)

type timeManager struct {
	db *pgx.Conn
}

func NewTime(db *pgx.Conn) *timeManager {
	return &timeManager{
		db: db,
	}
}

func (t *timeManager) CreateTime(ctx context.Context, user *dto.CreateUser) error {
	query := "INSERT INTO time (USER_TG_ID, WAKEUP_TIME, GOING_TO_BED) VALUES ($1, $2, $3)"
	if _, err := t.db.Exec(ctx, query, user.TgID, user.WakeupTime, user.GoingToBedTime); err != nil {
		return fmt.Errorf("CreateTime: failed insert new record to time table: %w", err)
	}

	return nil
}

func (t *timeManager) GetGoingToBedTime(ctx context.Context, tgID string) (*time.Time, error) {
	var userTime time.Time
	query := "SELECT GOING_TO_BED FROM time WHERE USER_TG_ID = $1"

	row := t.db.QueryRow(ctx, query, tgID)
	if err := row.Scan(&userTime); err != nil {
		return nil, fmt.Errorf("GetGoingToBedTime: failed get user going to bed time: %w", err)
	}

	return &userTime, nil
}

func (t *timeManager) GetWakeupTime(ctx context.Context, tgID string) (*time.Time, error) {
	var userTime time.Time
	query := "SELECT WAKEUP_TIME FROM time WHERE USER_TG_ID = $1"

	row := t.db.QueryRow(ctx, query, tgID)
	if err := row.Scan(&userTime); err != nil {
		return nil, fmt.Errorf("GetWakeupTime: failed get user wakeup time: %w", err)
	}

	return &userTime, nil
}

func (t *timeManager) GetTime(ctx context.Context, tgID string) (*models.Time, error) {
	var time models.Time
	query := "SELECT * FROM time WHERE USER_TG_ID = $1"

	row := t.db.QueryRow(ctx, query, tgID)
	if err := row.Scan(&time); err != nil {
		return nil, fmt.Errorf("GetTime: failed get user time: %w", err)
	}

	return &time, nil
}

func (t *timeManager) UpdateWakeupTime(ctx context.Context, time dto.UpdateWakeupTime) error {
	query := "UPDATE time SET WAKEUP_TIME = $1 WHERE USER_TG_ID = $2"

	if _, err := t.db.Exec(ctx, query, time.WakeupTime, time.TgID); err != nil {
		return fmt.Errorf("UpdateWakeupTime: failed update user wakeup time: %w", err)
	}

	return nil
}

func (t *timeManager) UpdateGoingToBedTime(ctx context.Context, time dto.UpdateGoingToBedTime) error {
	query := "UPDATE time SET GOING_TO_BED = $1 WHERE USER_TG_ID = $2"

	if _, err := t.db.Exec(ctx, query, time.GoingToBedTime, time.TgID); err != nil {
		return fmt.Errorf("UpdateGoingToBedTime: failed update user going to bed time: %w", err)
	}

	return nil
}
