package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin/json"
	"match-server/model"
	"match-server/utils"
)

func Login_in(ctx *gin.Context) {
	//日志
	logger := utils.Logger

	//用户登陆
	uid_cookie:=&http.Cookie{
		Name:   "admin",
		Value:   "123456",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer,uid_cookie)
	uid := ctx.Query("uid")
	//TODO mongo 持久化
	DB := utils.ShowDB()
	defer DB.Session.Close()
	coll := DB.C("user")
	user := model.User{}
	err_qu := coll.FindId(bson.ObjectIdHex(uid)).One(&user)

	if len(user.Id_) ==0 || err_qu != nil {
		logger.Info("查询日志", err_qu)
		ctx.JSON(http.StatusOK, gin.H{"message": "用户未注册"})
		return
	}
	//转化json
	jsons,_ :=json.Marshal(user)
	_,err2 := utils.Client.Do("SET","user_"+user.Id_.Hex(),jsons,"EX",45*60).Result()
	if err2 != nil {logger.Info(err2)}
	ctx.JSON(http.StatusOK,"OK")
}
