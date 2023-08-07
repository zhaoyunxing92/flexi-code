package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/service/datasheet"
)

type DatasheetCtrl struct {
	datasheet *datasheet.Service
}

func NewDatasheetCtrl(datasheet *datasheet.Service) *DatasheetCtrl {
	return &DatasheetCtrl{datasheet: datasheet}
}

func (ctrl *DatasheetCtrl) Save(ctx *gin.Context) {

}
