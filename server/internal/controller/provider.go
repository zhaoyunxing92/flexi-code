package controller

import (
	"github.com/google/wire"
)

var ProviderSetCtrl = wire.NewSet(
	NewDatasheetCtrl,
	NewAccountCtrl,
)
