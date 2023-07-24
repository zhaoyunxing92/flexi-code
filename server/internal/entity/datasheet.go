package entity

import (
	"gorm.io/gorm"
)

type Datasheet struct {
	// tenant code
	TenantId string `gorm:"index;size:32;not null"`

	// app code
	AppId string `gorm:"index;size:32;not null"`

	// module code
	ModuleId string `gorm:"index;size:32"`

	// datasheet code
	Code string `gorm:"uniqueIndex;size:32;not null"`

	// datasheet name
	Name string `gorm:"size:32;not null"`

	// datasheet icon
	Icon string `gorm:"size:16"`

	// datasheet desc
	Desc string `gorm:"size:256"`

	// datasheet alias
	Alias string `gorm:"size:16"`

	// datasheet version
	Version uint `gorm:"size:8;default:1"`

	gorm.Model
}
