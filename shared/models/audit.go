package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Audit struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt sql.NullTime   `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}
