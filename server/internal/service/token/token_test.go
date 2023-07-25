package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/sys"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

func TestGenerateToken(t *testing.T) {
	log.NewDefault()

	service := NewTokenService(&configs.App{Token: configs.Token{Secret: "chat", Expired: 30}})
	token, err := service.GenerateToken(sys.Account{Name: "fc"})

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	t.Log(token)

	claims, err := service.ParseToken(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, claims)
}
