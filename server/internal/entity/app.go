package entity

import (
	"gorm.io/gorm"
)

type App struct {
	// tenant code
	TenantId string `gorm:"index;size:32;not null"`

	// app code
	Code string `gorm:"uniqueIndex;size:32;not null"`

	// app name
	Name string `gorm:"size:16;not null"`

	gorm.Model
}
