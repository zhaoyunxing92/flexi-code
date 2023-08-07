package sys

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo/sys"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
)

type AccountService struct {
	repo sys.AccountRepo
}

func NewAccountService(repo sys.AccountRepo) *AccountService {
	return &AccountService{repo: repo}
}

// Login email and password login
func (s *AccountService) Login(email, password string) (schema.LoginResp, error) {
	return s.repo.Login(email, password)
}
