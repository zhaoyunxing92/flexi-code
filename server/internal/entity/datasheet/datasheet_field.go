package datasheet

import (
	"database/sql"

	"gorm.io/gorm"
)

type DatasheetField struct {

	// table code
	Table string `gorm:"index;size:32;not null"`

	// field code
	Code string `gorm:"index;size:32;not null"`

	// field name
	Name string `gorm:"size:32;not null"`

	// field alias
	Alias string `gorm:"size:32"`

	// field icon
	Icon string `gorm:"size:16"`

	// field placeholder
	Placeholder string `gorm:"size:32"`

	// field type
	Type string `gorm:"size:16;not null"`

	// field Required
	Required sql.NullBool `gorm:"not null;default:false"`

	// field disabled
	Disabled sql.NullBool `gorm:"not null;default:false"`

	// field order
	Order float32 `gorm:"not null;default:0"`

	// field default value
	DefaultValue string `gorm:"size:32"`

	// field format
	Format string `gorm:"size:16"`

	// field options
	Options []DatasheetFieldOption `gorm:"type:json"`

	// field desc
	Desc string `gorm:"size:128"`

	gorm.Model
}

type DatasheetFieldOption struct {
	// option label
	Label string `json:"label"`

	// option value
	Value string `json:"value"`
}
