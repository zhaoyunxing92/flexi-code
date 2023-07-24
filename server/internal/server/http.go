package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/internal/router"
)

func NewHttpServer(router *router.Router, app *configs.App) *gin.Engine {

	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.GET("/health", func(ctx *gin.Context) { ctx.String(200, "OK") })

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"version": app.Version,
			"title":   app.Name,
		})
	})

	v1 := r.Group("/api/v1")
	{
		datasheet := v1.Group("/datasheet")
		router.RegisterDatasheet(datasheet)
	}
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	//r.GET("/health", funcs(ctx *gin.Context) { ctx.String(200, "OK") })
	//
	//login := r.Group("/api/v1")
	//router.RegisterRouter(login)
	//
	//chat := r.Group("/api/v1/chat")
	//chat.Use(middleware.Authorized())
	//router.RegisterChatRouter(chat)

	return r
}
