package service

import (
	"context"
	"encoding/json"
	"github.com/cihub/seelog"
	"match-server/model"
	"match-server/utils"
)

var (
	bidOrder,askOrder *model.Order
	bidPo,askPo *model.Position
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


//校验资产和订单数据
func (t *Trade)CheckTrade() bool {
	//获取买单
	bidOrder = t.searchOrderById(t.BidId)
	//获取卖单
	askOrder= t.searchOrderById(t.AskId)
	//获取保证金
	bidBalance,err1 := SearchBalance(t.BidUid,accountType[UMargin]["BTC"],false)
	if err1!=nil {
		return false
	}
	askBalance,err2 := SearchBalance(t.BidUid,accountType[UMargin]["BTC"],false)
	if err2!=nil {
		return false
	}
	//校验订单和手续费账户余额
	if !t.chargeOrder(bidOrder,bidBalance) || !t.chargeOrder(askOrder,askBalance) {
		return false
	}

	return true
}

//划转资产，完成数据持久化
func (t *Trade)TransferAssets() error {
	//1.手续费流水
	var bidTrans = t.tradeFee(bidOrder)
	var askTrans = t.tradeFee(askOrder)
	//2.计算平仓收益
	bidPo = t.searchPo(bidOrder)
	askPo = t.searchPo(askOrder)

	bidCloTrans := t.closePo(bidOrder,bidPo)
	askCloTrans := t.closePo(askOrder,askPo)
	//3.爆仓单划转收益到风险准备金


	err2 := insertTrans(bidTrans,askTrans,bidCloTrans,askCloTrans)
	return err2
}