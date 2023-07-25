package space

import (
	"gorm.io/gorm"
)

type Space struct {
	// space id
	SpaceId string `gorm:"uniqueIndex;size:32;not null"`

	// space name
	Name string `gorm:"size:32;not null"`

	gorm.Model
}
