package ctrl

import (
	"github.com/google/wire"
)

var ProviderSetCtrl = wire.NewSet(
	NewDatasheetCtrl,
)
