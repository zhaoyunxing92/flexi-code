package router

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/ctrl"
)

type Router struct {
	datasheet *ctrl.DatasheetCtrl
}

func NewRouter(datasheet *ctrl.DatasheetCtrl) *Router {
	return &Router{datasheet: datasheet}
}

func (api *Router) RegisterDatasheet(r *gin.RouterGroup) {
	r.POST("/save", api.datasheet.Save)
}
