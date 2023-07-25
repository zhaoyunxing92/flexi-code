package service

import (
	"github.com/google/wire"

	"github.com/zhaoyunxing92/flexi-code/server/internal/service/datasheet"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service/sys"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service/token"
)

var ProviderSetService = wire.NewSet(
	token.NewTokenService,
	sys.NewAccountService,
	datasheet.NewDatasheetService,
)
