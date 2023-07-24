package service

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
)

type AccountService struct {
	repo repo.AccountRepo
}

func NewAccountService(repo repo.AccountRepo) *AccountService {
	return &AccountService{repo: repo}
}

// Login email and password login
func (s *AccountService) Login(email, password string) (schema.LoginResp, error) {
	return s.repo.Login(email, password)
}
