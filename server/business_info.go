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
	BidOrder model.Order
 	AskOrder model.Order
 	BidBalance float64
 	AskBalance float64
)

func (t *Trade)CheckTrade() bool {
	var err error
	BidOrder,err = t.SearchOrderById(t.BidId)
	if err != nil {
		panic(err)
	}
	AskOrder,err = t.SearchOrderById(t.AskId)
	if err != nil {
		panic(err)
	}
	t.chargeFee(&AskOrder)
	t.chargeFee(&BidOrder)

	return true
}

func (t *Trade)SearchOrderById(id uint) (model.Order, error) {
	var order model.Order
	sql := "select * from co_order_"+ strings.ToLower(t.Symbol) +" where id = ?"
	err := utils.DBContract().SQL(sql, id).Find(&order)
	if err != nil {
		seelog.Error("order not exists",err)
		return order,err
	}
	return order,nil
}

func (t *Trade)chargeFee(order *model.Order) float64 {
	var feeRate = order.FeeRateMaker
	if t.TrendSide == order.Side {
		feeRate = order.FeeRateTaker
	}
	return float64(t.Volume) * t.Price * feeRate
}