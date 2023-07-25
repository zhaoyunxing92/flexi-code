package datasheet

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo/datasheet"
)

type Service struct {
	repo datasheet.Repo
}

func NewDatasheetService(repo datasheet.Repo) *Service {
	return &Service{repo: repo}
}
