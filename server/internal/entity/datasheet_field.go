package entity

import (
	"gorm.io/gorm"
)

type DatasheetField struct {

	// table code
	Table string

	// field code
	Code string

	// field name
	Name string

	// field alias
	Alias string

	// field icon
	Icon string

	// field placeholder
	Placeholder string

	// field type
	Type string

	// field Required
	Required bool

	// field disabled
	Disabled bool

	// field order
	Order string

	// field default value
	DefaultValue string

	// field format
	Format string

	// field options
	Options []DatasheetFieldOption

	// field desc
	Desc string

	gorm.Model
}

func (DatasheetField) TableName() string {
	return "datasheet_field"
}

type DatasheetFieldOption struct {
	// option label
	Label string

	// option value
	Value string
}
