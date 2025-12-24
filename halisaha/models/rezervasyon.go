package models

import "time"

type Rezervasyon struct {
	ID     uint      `json:"id" grom:"primaryKey" `
	UserID uint      `json:"user_id"`
	SahaID uint      `json:"saha_id"`
	Tarih  time.Time `json:"tarih"`
	Saat   string    `json:"saat"`
}
