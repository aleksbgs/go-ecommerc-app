package domain

import "time"

type Address struct {
	ID           uint      `gorm:"primary_key"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	PostCode     uint      `json:"post_code"`
	Country      string    `json:"country"`
	UserId       uint      `json:"user_id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
