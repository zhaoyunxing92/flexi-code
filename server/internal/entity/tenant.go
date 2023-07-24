package entity

import (
	"gorm.io/gorm"
)

type Tenant struct {
	// tenant code
	Code string `gorm:"uniqueIndex;size:32;not null"`

	// tenant name
	Name string `gorm:"size:32;not null"`

	gorm.Model
}
