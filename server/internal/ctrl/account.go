package ctrl

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/handler"
	"github.com/zhaoyunxing92/flexi-code/server/internal/models"
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
	if res, err := ctrl.accountService.Login(req.Email, req.Password); err != nil {
		handler.Response(ctx, models.Fail(400, "失败", err.Error()))
	} else {
		handler.Response(ctx, models.Succeed("登录成功", res))
	}
}
