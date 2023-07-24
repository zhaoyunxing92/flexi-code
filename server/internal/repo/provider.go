package repo

import (
	"github.com/google/wire"
)

var ProviderSetRepo = wire.NewSet(
	NewAccountRepo,
	NewDatasheetRepo,
)
