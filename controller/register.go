package controller

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"match-server/model"
	"match-server/utils"
)

func Register(ctx *gin.Context) {
	//用户注册
	uid := ctx.Query("uid")
	//TODO mongo 持久化
	DB := utils.ShowDB()
	defer DB.Session.Close()
	coll := DB.C("user")

	objectId := bson.NewObjectId()

	err1 := coll.Insert(&model.User{objectId,uid,25})
	if err1 !=nil {
		seelog.Info("持久化异常",err1)
	}
}
