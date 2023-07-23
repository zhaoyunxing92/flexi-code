package entity

import (
	"gorm.io/gorm"
)

type Tenant struct {
	// tenant code
	Code string

	// tenant name
	Name string

	gorm.Model
}

func (Tenant) TableName() string {
	return "tenant"
}
