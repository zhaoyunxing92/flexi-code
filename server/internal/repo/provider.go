package repo

import (
	"github.com/google/wire"

	"github.com/zhaoyunxing92/flexi-code/server/internal/repo/datasheet"
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo/sys"
)

var ProviderSetRepo = wire.NewSet(
	sys.NewAccountRepo,
	datasheet.NewDatasheetRepo,
	datasheet.NewFieldRepo,
)
