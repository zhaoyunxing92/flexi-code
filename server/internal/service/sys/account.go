package sys

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo/sys"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/errors"
)

type AccountService struct {
	repo sys.AccountRepo
}

func NewAccountService(repo sys.AccountRepo) *AccountService {
	return &AccountService{repo: repo}
}

// Login email and password login
func (s *AccountService) Login(email, password string) (schema.LoginResp, errors.Error) {
	return s.repo.Login(email, password)
}
