package models

type Setting struct {
	ID       uint         `gorm:"primaryKey" json:"id"`
	UserID   uint         `gorm:"not null;uniqueIndex:idx_user_setting" json:"userId"`
	Language LanguageType `gorm:"not null" json:"language"`
	Currency CurrencyType `gorm:"not null" json:"currency"`
}
