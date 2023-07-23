package storage

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

var tables = []interface{}{
	&entity.App{},
	&entity.Tenant{},
	&entity.Module{},
	&entity.Datasheet{},
	&entity.DatasheetField{},
	&entity.DatasheetRecord{},
}

type Storage struct {
	DB     *gorm.DB
	config *configs.Storage
}

func NewStorage(config *configs.Storage) (storage *Storage, err error) {
	storage = &Storage{config: config}
	return storage, storage.InitStorageTable()
}

func (storage *Storage) InitStorageTable() error {
	var (
		err    error
		db     *gorm.DB
		config = storage.config
		conf   = &gorm.Config{}
		driver gorm.Dialector
	)
	if config.IsDebug() {
		conf.Logger = logger.Default.LogMode(logger.Info)
	}
	dsn := config.Connection
	switch config.Driver {
	default:
		driver = mysql.Open(dsn)
	}

	// table prefix
	conf.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   config.TablePrefix,
		SingularTable: true,
	}

	if db, err = gorm.Open(driver, conf); err != nil {
		return err
	}

	if config.IsAutoMigrate() {
		if err = db.AutoMigrate(tables...); err != nil {
			log.Errorf("auto migrate tables failed: %s", err)
			return err
		}
	}

	storage.DB = db
	return storage.Ping()
}

func (storage *Storage) CreateDirIfNotExist(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (storage *Storage) Ping() error {
	if db, err := storage.DB.DB(); err != nil {
		return err
	} else {
		return db.Ping()
	}
}
