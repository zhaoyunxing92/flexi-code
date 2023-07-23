package entity

import (
	"gorm.io/gorm"
)

type DatasheetRecord struct {
	// tenant code
	TenantId string

	// app code
	AppId string

	// module code
	ModuleId string

	// table code
	Code string

	// table name
	Name string

	// table icon
	Icon string

	// table desc
	Desc string

	// table alias
	Alias string

	// table version
	Version uint

	gorm.Model
}

func (DatasheetRecord) TableName() string {
	return "datasheet_record"
}
