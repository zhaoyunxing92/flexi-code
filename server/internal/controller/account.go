package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/handler"
	"github.com/zhaoyunxing92/flexi-code/server/internal/schema"
	"github.com/zhaoyunxing92/flexi-code/server/internal/service/sys"
)

type AccountCtrl struct {
	accountService *sys.AccountService
}

func NewAccountCtrl(accountService *sys.AccountService) *AccountCtrl {
	return &AccountCtrl{accountService: accountService}
}

func (ctrl *AccountCtrl) Login(ctx *gin.Context) {
	req := schema.LoginReq{}
	if handler.BindAndCheck(ctx, &req) {
		return
	}
	resp, err := ctrl.accountService.Login(req.Email, req.Password)
	handler.Response(ctx, err, resp)
}
