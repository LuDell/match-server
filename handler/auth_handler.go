package handler

import (
	"github.com/gin-gonic/gin"
	"match-server/utils"
)

func AuthHandler(ctx *gin.Context)  {
	token := ctx.GetHeader("token")
	val,err := utils.Client.Do("GET","user_"+token).Result()
	if err != nil || val == nil {
		ctx.Abort()
	}
	ctx.Set("landing",true)

	ctx.Next()
}
