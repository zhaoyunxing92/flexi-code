package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	NickName string `gorm:"size:64"`

	Password string `gorm:"size:128;not null"`

	Email string `gorm:"uniqueIndex;size:128;not null"`

	Gender int `gorm:"size:1"`

	// account avatar
	Avatar string `gorm:"size:256"`

	Locale string `gorm:"size:32"`

	TimeZone string `gorm:"size:32"`

	gorm.Model
}
