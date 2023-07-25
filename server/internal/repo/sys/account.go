package sys

import (
	"gorm.io/gorm"

	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/sys"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service/token"
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/errors"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/tools"
)

type AccountRepo interface {
	Login(email, password string) (schema.LoginResp, errors.Error)
}

type account struct {
	db           *gorm.DB
	tokenService *token.Service
}

func NewAccountRepo(storage *storage.Storage, tokenService *token.Service) AccountRepo {
	return &account{db: storage.DB, tokenService: tokenService}
}

func (a *account) Login(email, pwd string) (schema.LoginResp, errors.Error) {
	var (
		err  error
		acc  sys.Account
		res  schema.LoginResp
		jwt  string
		hash string
	)

	acc = sys.Account{}
	if err = a.db.FirstOrCreate(&acc, sys.Account{Email: tools.ToNullString(email)}).Error; err != nil {
		return schema.LoginResp{}, errors.New(err.Error())
	}

	if hash, err = tools.GeneratePassword(pwd); err != nil {
		return schema.LoginResp{}, errors.New("系统异常")
	}

	if len(acc.Password) == 0 {
		acc.Password = hash
		if err = a.db.Save(acc).Error; err != nil {
			return schema.LoginResp{}, errors.New("系统异常")
		}
	} else if !tools.ComparePassword(hash, pwd) {
		return schema.LoginResp{}, errors.New("账号或密码错误")
	}

	if jwt, err = a.tokenService.GenerateToken(acc); err != nil {
		return schema.LoginResp{}, errors.New(err.Error())
	}

	res.Avatar = acc.Avatar
	res.Name = acc.Name
	res.Token = jwt
	return res, nil
}
