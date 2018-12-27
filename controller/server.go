package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"match-server/utils"
	"time"
)

var(
	upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func WsHandler( ctx *gin.Context)  {

	w := ctx.Writer
	r := ctx.Request

	var (
		wsConn *websocket.Conn
		err error
		data []byte
		conn *utils.Connection
	)
	if wsConn,err = upgrader.Upgrade(w,r,nil); err != nil{
		return
	}

	conn = utils.InitConnection(wsConn)

	go func() {
		for {
			if err := conn.WriteMessage([]byte("heartbeat"));err != nil{
				return
			}
			time.Sleep(90 *time.Second)
		}
	}()

	for {
		if data,err = conn.ReadMessage();err != nil {
			goto ERROR
		}

		if err = conn.WriteMessage(data);err != nil {
			goto ERROR
		}
	}
ERROR:
	conn.Close()
}
