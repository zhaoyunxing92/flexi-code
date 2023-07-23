package entity

import (
	"gorm.io/gorm"
)

type Module struct {
	// tenant code
	TenantId string

	// app code
	AppId string

	// module code
	Code string

	// module name
	Name string

	// module icon
	Icon string

	// module desc
	Desc string

	// module order
	Order uint

	gorm.Model
}

func (Module) TableName() string {
	return "module"
}
