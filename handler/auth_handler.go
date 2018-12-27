package handler

import (
	"github.com/gin-gonic/gin"
	"match-server/utils"
)

func AuthHandler(ctx *gin.Context)  {
	token := ctx.GetHeader("token")
	conn := utils.RedisPool.Get()
	defer conn.Close()
	val,err := conn.Do("GET","user_"+token)
	if err != nil || val == nil {
		ctx.Abort()
	}
	ctx.Set("landing",true)

	ctx.Next()
}
