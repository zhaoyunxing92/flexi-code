package internal

import (
	"github.com/google/wire"

	"github.com/zhaoyunxing92/flexi-code/server/internal/router"
	"github.com/zhaoyunxing92/flexi-code/server/internal/server"
	"github.com/zhaoyunxing92/flexi-code/server/internal/storage"
)

var ProviderSetBase = wire.NewSet(
	storage.NewStorage,
	router.NewRouter,
	server.NewHttpServer,
)
