package repo

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

type DatasheetRepo interface {
	Save() error
}

type datasheet struct {
	storage *storage.Storage
}

func NewDatasheetRepo(storage *storage.Storage) DatasheetRepo {
	return &datasheet{storage: storage}
}

func (d *datasheet) Save() error {
	return nil
}
