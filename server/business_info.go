package server

import (
	"github.com/cihub/seelog"
	"github.com/go-xorm/builder"
	"match-server/model"
	"match-server/utils"
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

var err error

var (
	BidOrder model.Order
 	AskOrder model.Order
 	BidBalance float64
 	AskBalance float64
)

func (t *Trade)CheckTrade() bool {
	BidOrder,err = SearchOrderById(t.BidId)
	if err != nil {
		panic(err)
	}
	AskOrder,err = SearchOrderById(t.AskId)
	if err != nil {
		panic(err)
	}
	t.chargeFee(&AskOrder)
	t.chargeFee(&BidOrder)

	return true
}

func SearchOrderById(id uint) (model.Order, error) {
	var order model.Order
	err := utils.DBContract().Where(builder.Eq{"id":id}).Find(&order)
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