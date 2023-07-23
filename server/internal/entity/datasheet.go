package entity

import (
	"gorm.io/gorm"
)

type Datasheet struct {
	// tenant code
	TenantId string

	// app code
	AppId string

	// module code
	ModuleId string

	// datasheet code
	Code string

	// datasheet name
	Name string

	// datasheet icon
	Icon string

	// datasheet desc
	Desc string

	// datasheet alias
	Alias string

	// datasheet version
	Version uint

	gorm.Model
}

func (Datasheet) TableName() string {
	return "datasheet"
}
