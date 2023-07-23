package ctrl

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/service"
)

type DatasheetCtrl struct {
	datasheet *service.DatasheetService
}

func NewDatasheetCtrl(datasheet *service.DatasheetService) *DatasheetCtrl {
	return &DatasheetCtrl{datasheet: datasheet}
}

func (ctrl *DatasheetCtrl) Save(ctx *gin.Context) {

}
