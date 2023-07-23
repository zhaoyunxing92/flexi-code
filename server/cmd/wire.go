//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/internal"
	"github.com/zhaoyunxing92/flexi-code/server/internal/ctrl"
	"github.com/zhaoyunxing92/flexi-code/server/internal/repo"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service"
)

// initApplication init application.
func initApplication(appConfig *configs.App, storageConfig *configs.Storage) (*gin.Engine, error) {
	panic(wire.Build(
		internal.ProviderSetBase,
		repo.ProviderSetRepo,
		ctrl.ProviderSetCtrl,
		service.ProviderSetService,
	))
}
