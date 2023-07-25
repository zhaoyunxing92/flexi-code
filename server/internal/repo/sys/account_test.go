package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhaoyunxing92/flexi-code/server/internal/config"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service/token"
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

func TestLogin(t *testing.T) {
	conf, err := config.Load("../../../conf/application.yaml")
	assert.Nil(t, err)

	db, err := storage.NewStorage(conf.Storage)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	tokenService := token.NewTokenService(conf.App)
	assert.Nil(t, err)

	res, err := NewAccountRepo(db, tokenService).
		Login("zhaoyunxing@apache.org", "12345678")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
