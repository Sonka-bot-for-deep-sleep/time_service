package models

import "time"

type Time struct {
	ID         int       `json:"id"`
	UserTgID   string    `json:"user_tg_id" sql:"USER_TG_ID"`
	WakeupTime time.Time `json:"wakeup_time" sql:"WAKEUP_TIME"`
	GoingToBed time.Time `json:"going_to_bed" sql:"GOING_TO_BED"`
}
