package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/internal/models"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

// Response handle response body
func Response(ctx *gin.Context, err error, data models.Response[models.Body]) {
	//lang := GetLang(ctx)
	// no error
	if err == nil {
		ctx.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, reason.Success, data).TrMsg(lang))
		return
	}

	// no error
	ctx.JSON(http.StatusOK, data)
}

// BindAndCheck bind request and check
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Errorf("参数绑定异常: %s", err.Error())
		Response(ctx, models.Fail(400, "参数绑定异常", "参数绑定异常"))
		ctx.Abort()
		return true
	}

	//errField, err := validator.GetValidatorByLang(lang).Check(data)
	//if err != nil {
	//	HandleResponse(ctx, err, errField)
	//	return true
	//}
	return false
}
