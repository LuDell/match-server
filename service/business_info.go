package service

import (
	"github.com/cihub/seelog"
	"match-server/model"
	"match-server/utils"
	"math"
	"strings"
	"time"
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
	bidOrder *model.Order
 	askOrder *model.Order
 	bidBalance float64
 	askBalance float64
)

func (t *Trade)CheckTrade() bool {
	//获取买单
	bidOrder = t.searchOrderById(t.BidId)
	//获取卖单
	askOrder= t.searchOrderById(t.AskId)
	//获取保证金
	bidBalance,err1 := SearchBalance(t.BidUid,accountType[U_MARGIN]["BTC"],false)
	if err1!=nil {
		return false
	}
	askBalance,err2 := SearchBalance(t.BidUid,accountType[U_MARGIN]["BTC"],false)
	if err2!=nil {
		return false
	}
	//校验订单和手续费账户余额
	if !t.chargeOrder(bidOrder,bidBalance) {
		return false
	}
	if !t.chargeOrder(askOrder,askBalance) {
		return false
	}
	return true
}

func (t *Trade)chargeOrder(order *model.Order,balance float64) bool {
	//订单未成交数量和当前trade
	var voIsOk = t.Volume <= order.Volume-order.DealVolume
	//交易手续费
	exchangeFee,_ := t.transferFee(order)

	return exchangeFee <= balance && voIsOk
}

func (t *Trade)transferFee(order *model.Order) (float64,bool) {
	var feeRate = order.FeeRateMaker
	var isTaker = false
	if t.TrendSide == order.Side {
		feeRate = order.FeeRateTaker
		isTaker = true
	}
	var exchangeFee = float64(t.Volume) * t.Price * feeRate
	return exchangeFee,isTaker
}

//查询订单
func (t *Trade)searchOrderById(id uint) *model.Order {
	var order model.Order
	sql := "select * from co_order_"+ strings.ToLower(t.Symbol) +" where id = ?"
	if err := utils.DBContract().SQL(sql, id).Find(&order); err != nil {
		seelog.Error("order not exists",err)
		return &order
	}
	return &order
}

//划转资产
func (t *Trade)TransferAssets() error {
	//获取系统手续费账户
	sysFeeBalance,err1 := SearchBalance(sysUid,accountType[C_EXCHANGE_FEE][strings.ToUpper(t.Symbol)],true)
	if err1 != nil {
		return err1
	}
	//1.手续费流水
	var bidTrans = t.tradeFee(bidOrder,bidBalance,sysFeeBalance)
	var askTrans = t.tradeFee(bidOrder,bidBalance,sysFeeBalance)
	//TODO
	//2.计算平仓收益

	//3.爆仓单风险准备

	_,err2 := insertTrans(bidTrans,askTrans)
	return err2
}


//获取用户仓位
func (t *Trade)tradeFee(order *model.Order,marginBalance float64,sysFeeBalance float64) *model.Transaction{
	orderFee,isTaker := t.transferFee(order)
	var feeTrans = model.Transaction{}
	if orderFee >0 {
		feeTrans.FromUid = order.Uid
		feeTrans.FromType = accountType[U_MARGIN][strings.ToUpper(t.Symbol)]
		feeTrans.FromBalance = marginBalance - math.Abs(orderFee)
		feeTrans.ToUid = sysUid
		feeTrans.ToType = accountType[C_EXCHANGE_FEE][strings.ToUpper(t.Symbol)]
		feeTrans.ToBalance = sysFeeBalance + math.Abs(orderFee)
	}else{
		feeTrans.ToUid = order.Uid
		feeTrans.ToType = accountType[U_MARGIN][strings.ToUpper(t.Symbol)]
		feeTrans.ToBalance = marginBalance - math.Abs(orderFee)
		feeTrans.FromUid = sysUid
		feeTrans.FromType = accountType[C_EXCHANGE_FEE][strings.ToUpper(t.Symbol)]
		feeTrans.FromBalance = sysFeeBalance + math.Abs(orderFee)
	}
	feeTrans.Amount = math.Abs(orderFee)
	feeTrans.Meta = "交易手续费"
	feeTrans.Scene = FEE_MAKER
	if isTaker {feeTrans.Scene = FEE_TAKER}
	feeTrans.RefType = t.Symbol
	feeTrans.RefId = order.Id
	feeTrans.Op_uid = 0
	feeTrans.Op_ip = "0.0.0.0"
	feeTrans.Ctime = time.Now().Unix()
	feeTrans.Mtime = time.Now().Unix()
	return &feeTrans
}