package service

import (
	"context"
	"encoding/json"
	"github.com/cihub/seelog"
	"match-server/utils"
)

type Data struct {
	code int
	msg string
	data map[string]string
}

type Persist int

func (p *Persist)DataPersistence(ctx context.Context, args []byte, reply *[]byte) error {
	var trade Trade
	err := json.Unmarshal(args,&trade)
	if err != nil {
		seelog.Error("trade unmarshal is error")
		return err;
	}

	bool := trade.CheckTrade()
	if !bool {
		var resData = Data{1001,"订单或者账户异常",nil}
		response,_ := json.Marshal(resData)
		reply = &response
	}
	//开始事物
	session := utils.DBExchange().NewSession()
	session.Begin()

	return nil
}