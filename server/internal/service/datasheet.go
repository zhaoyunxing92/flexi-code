package service

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo"
)

type DatasheetService struct {
	repo repo.DatasheetRepo
}

func NewDatasheetService(repo repo.DatasheetRepo) *DatasheetService {
	return &DatasheetService{repo: repo}
}
