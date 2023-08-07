package router

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/controller"
)

type Router struct {
	datasheet *controller.DatasheetCtrl
}

func NewRouter(datasheet *controller.DatasheetCtrl) *Router {
	return &Router{datasheet: datasheet}
}

func (api *Router) RegisterDatasheet(r *gin.RouterGroup) {
	r.POST("/save", api.datasheet.Save)
}
