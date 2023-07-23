package entity

import (
	"gorm.io/gorm"
)

type App struct {
	// tenant code
	TenantId string `gorm:"index;size:32;not null"`

	// app code
	Code string

	// app name
	Name string

	gorm.Model
}

func (App) TableName() string {
	return "app"
}
