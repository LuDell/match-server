package server

import (
	"github.com/cihub/seelog"
	"match-server/model"
	"match-server/utils"
	"strings"
)

type Trade struct {
	Symbol string
	ContractId uint
	Price float64
	Volume uint
	BidId uint
	AskId uint
	TrendSide string
	BidUid uint
	AskUid uint
	BidFromStatus uint8
	BidToStatus uint8
	AskToStatus uint8
	Ctime int
	Token string
}


var (
	BidOrder *model.Order
 	AskOrder *model.Order
 	BidBalance float64
 	AskBalance float64
)

func (t *Trade)CheckTrade() bool {
	//获取买单
	BidOrder = t.SearchOrderById(t.BidId)
	//获取卖单
	AskOrder= t.SearchOrderById(t.AskId)

	t.chargeOrder(AskOrder)
	t.chargeOrder(BidOrder)

	return true
}

func (t *Trade)SearchOrderById(id uint) *model.Order {
	var order model.Order
	sql := "select * from co_order_"+ strings.ToLower(t.Symbol) +" where id = ?"
	if err := utils.DBContract().SQL(sql, id).Find(&order); err != nil {
		seelog.Error("order not exists",err)
		return &order
	}
	return &order
}

func (t *Trade)chargeOrder(order *model.Order) (float64,bool) {
	var feeRate = order.FeeRateMaker
	if t.TrendSide == order.Side {
		feeRate = order.FeeRateTaker
	}
	var voIsOk = t.Volume <= order.Volume-order.DealVolume
	return float64(t.Volume) * t.Price * feeRate, voIsOk
}