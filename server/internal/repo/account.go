package repo

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/zhaoyunxing92/flexi-code/server/internal/entity"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service"
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

type AccountRepo interface {
	Login(email, password string) (schema.LoginResp, error)
}

type account struct {
	db    *gorm.DB
	token *service.TokenService
}

func NewAccountRepo(storage *storage.Storage, token *service.TokenService) AccountRepo {
	return &account{db: storage.DB, token: token}
}

func (a *account) Login(email, pwd string) (res schema.LoginResp, err error) {

	var user entity.Account
	if err = a.db.Table("account").Where("email = ?", email).First(&user).Error; err != nil {
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err = bcrypt.CompareHashAndPassword(hash, []byte(user.Password)); err != nil {
		return schema.LoginResp{}, errors.New("账号或密码错误")
	}
	res.Avatar = user.Avatar
	res.Name = user.Name
	return res, nil
}
