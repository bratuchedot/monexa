package models

import (
	"database/sql"
	shared "monexa/shared/models"
	"time"
)

type Record struct {
	ID uint `gorm:"primaryKey" json:"id"`
	shared.Audit
	UserID          uint           `gorm:"not null;index" json:"userId"`
	CategoryID      uint           `gorm:"not null;index" json:"categoryId"`
	PaymentMethodID uint           `gorm:"not null;index" json:"paymentMethodId"`
	Type            string         `gorm:"not null" json:"type"`
	Amount          float64        `gorm:"not null" json:"amount"`
	Currency        CurrencyType   `gorm:"not null" json:"currency"`
	Description     sql.NullString `json:"description"`
	Date            time.Time      `gorm:"not null" json:"date"`
}
