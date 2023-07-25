package datasheet

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

type Repo interface {
	Save() error
}

type datasheet struct {
	storage *storage.Storage
}

func NewDatasheetRepo(storage *storage.Storage) Repo {
	return &datasheet{storage: storage}
}

func (d *datasheet) Save() error {
	return nil
}
