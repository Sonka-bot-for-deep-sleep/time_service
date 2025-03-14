package dto

type GetWakeupTime struct {
	TgID string `json:"tg_id"`
}

type GetGoingToBedTime struct {
	TgID string `json:"tg_id"`
}

type SetGoingToBedTime struct {
	TgID           string `json:"tg_id"`
	GoingToBedTime string `json:"going_to_bed_time"`
}

type SetWakeupTime struct {
	TgID       string `json:"tg_id"`
	WakeupTime string `json:"wakeup_time"`
}

type GetTime struct {
	TgID string `json:"tg_id"`
}

type UpdateWakeupTime struct {
	TgID       string `json:"tg_id"`
	WakeupTime string `json:"wakeup_time"`
}

type UpdateGoingToBedTime struct {
	TgID           string `json:"tg_id"`
	GoingToBedTime string `json:"going_to_bed_time"`
}

type CreateUser struct {
	TgID           string `json:"tg_id"`
	GoingToBedTime string `json:"going_to_bed_time"`
	WakeupTime     string `json:"wakeup_time"`
}
