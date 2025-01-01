package domain

import "time"

type BankAccount struct {
	ID          uint      `json:"id" gorm:"PrimaryKey"`
	UserId      string    `json:"user_id"`
	BankAccount uint      `json:"bank_account" gorm:"unique:not null"`
	SwiftCode   string    `json:"swift_code"`
	PaymentType string    `json:"payment_type"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
