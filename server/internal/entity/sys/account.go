package sys

import (
	"database/sql"
	"strings"

	"gorm.io/gorm"

	"github.com/zhaoyunxing92/flexi-code/server/pkg/tools"
)

type Account struct {
	// app code
	AccountId sql.NullString `gorm:"uniqueIndex;size:32;not null"`

	Name string `gorm:"size:32"`

	Password string `gorm:"size:128;not null"`

	Email sql.NullString `gorm:"uniqueIndex;size:128;not null"`

	Gender int `gorm:"size:1"`

	// account avatar
	Avatar string `gorm:"size:256"`

	Locale string `gorm:"size:32"`

	TimeZone string `gorm:"size:32"`

	gorm.Model
}

func (a Account) BeforeCreate(tx *gorm.DB) error {
	email := a.Email.String
	if len(a.Name) == 0 && len(email) > 0 {
		atIndex := strings.Index(email, "@")
		tx.Statement.SetColumn("name", email[:atIndex])
	}
	if len(a.AccountId.String) == 0 {
		tx.Statement.SetColumn("account_id", tools.ToNullString(tools.CreateAccount()))
	}
	return nil
}
