package entity

import (
	"gorm.io/gorm"
)

type DatasheetRecord struct {
	// tenant code
	TenantId string `gorm:"index;size:32;not null"`

	// app code
	AppId string `gorm:"index;size:32;not null"`

	// module code
	ModuleId string `gorm:"index;size:32"`

	// table code
	RecordId string `gorm:"uniqueIndex;size:32;not null"`

	Data string `gorm:"type:json"`

	// table version
	Version uint

	gorm.Model
}
