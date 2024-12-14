package models

import shared "monexa/shared/models"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	shared.Audit
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
}
