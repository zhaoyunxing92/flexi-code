package datasheet

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

type FieldRepo interface {
	Save() error
}

type field struct {
	storage *storage.Storage
}

func NewFieldRepo(storage *storage.Storage) FieldRepo {
	return &field{storage: storage}
}

func (d *field) Save() error {
	return nil
}
