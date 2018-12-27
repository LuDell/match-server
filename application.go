package main

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
	"match-server/controller"
	"match-server/handler"
	"time"
)

func main()  {
	//初始init
	app := gin.New()
	app.Use(gin.Recovery())
	app.Static("/static","views/static")
	app.StaticFile("/favicon.ico","views/static/favicon.ico")
	app.LoadHTMLGlob("views/*.html")

	app.GET("/",func(ctx *gin.Context) {

		cookie,err := ctx.Request.Cookie("admin")
		var value *string
		if(err == nil){
			value = &cookie.Value
		}

		seelog.Info("返回的cookie值", value)
		ctx.HTML(http.StatusOK,"index.html",gin.H{"data":"hello world"})

	})

	app.GET("/ws", controller.WsHandler)

	group := app.Group("api",handler.AuthHandler)
	{
		group.OPTIONS("public_info", func(ctx *gin.Context) {

			val,bol := ctx.Get("landing")

			ctx.JSON(http.StatusOK,gin.H{"val":val,"bol":bol,"data":"hello world"})
		})

	}

	app.GET("login_in",controller.Login_in)

	app.GET("register",controller.Register)

	server := &http.Server{
		Addr: ":8088",
		Handler: app,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 20*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	server.ListenAndServe()
}