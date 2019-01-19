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

func (t *Trade)chargeOrder(order *model.Order,balance float64) bool {
	//订单未成交数量和当前trade
	var voIsOk = t.Volume <= order.Volume-order.DealVolume
	//交易手续费
	exchangeFee,_ := t.transferFee(order)

	return (exchangeFee <= balance) && voIsOk
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

//交易手续费
func (t *Trade)tradeFee(order *model.Order) *model.Transaction{
	orderFee,isTaker := t.transferFee(order)
	var feeTrans = model.Transaction{}
	if orderFee >0 {
		feeTrans.FromUid = order.Uid
		feeTrans.FromType = accountType[UMargin]["BTC"]
		feeTrans.ToUid = sysUid
		feeTrans.ToType = accountType[CExchangeFee][strings.ToUpper(t.Symbol)]
	}else{
		feeTrans.ToUid = order.Uid
		feeTrans.ToType = accountType[UMargin]["BTC"]
		feeTrans.FromUid = sysUid
		feeTrans.FromType = accountType[CExchangeFee][strings.ToUpper(t.Symbol)]
	}
	feeTrans.Amount = math.Abs(orderFee)
	feeTrans.Meta = "交易手续费"
	feeTrans.Scene = FeeMaker
	if isTaker {feeTrans.Scene = FeeTaker}
	feeTrans.RefType = t.Symbol
	feeTrans.RefId = order.Id
	feeTrans.Op_uid = 0
	feeTrans.Op_ip = "0.0.0.0"
	feeTrans.Ctime = time.Now().Unix()
	feeTrans.Mtime = time.Now().Unix()
	return &feeTrans
}

func (t *Trade)closePo(o *model.Order,p *model.Position) *model.Transaction {
	//判断平仓
	if !strings.EqualFold(o.Side,p.Side) {
		//获取平仓数量
		var closeVo = math.Min(float64(o.Volume), float64(p.Volume))
		//计算平仓收益 多仓 = (平仓价格 - 持仓均价) * 持仓数量 * 乘数
		var income = closeVo * (t.Price - p.AvgPrice)
		if income != 0 {
			nowTime := time.Now().Unix()
			var transaction = model.Transaction{}
			transaction.Amount = math.Abs(income)
			transaction.Scene = utils.If(p.Side == Buy, CloseBuyProfitLoss, CloseSellProfitLoss).(string)
			transaction.Meta = utils.If(p.Side == Buy, CloseBuyProfitLoss, CloseSellProfitLoss).(string)
			transaction.RefId = o.Id
			transaction.RefType = "co_order_"+strings.ToLower(t.Symbol)
			transaction.Op_ip = ""
			transaction.Op_uid = 0
			transaction.Ctime = nowTime
			transaction.Mtime = nowTime
			if income > 0 {
				transaction.FromUid = sysUid
				transaction.FromType = accountType[CPositionClose][strings.ToUpper(t.Symbol)]
				transaction.ToUid = o.Uid
				transaction.ToType = accountType[UMargin]["BTC"]
			}else {
				transaction.ToUid = sysUid
				transaction.ToType = accountType[CPositionClose][strings.ToUpper(t.Symbol)]
				transaction.FromUid = o.Uid
				transaction.FromType = accountType[UMargin]["BTC"]
			}
			return &transaction
		}
	}
	return nil
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
//查询用户仓位
func (t *Trade)searchPo(order *model.Order) *model.Position {
	var position model.Position
	var sql = "select * from co_position where uid = ? and contract_id = ? and origin_oid = ?"
	if err := utils.DBContract().SQL(sql, order.Uid,order.ContractId,order.OriginOid).Find(&position); err != nil {
		seelog.Error("position is null",err)
		return nil
	}
	return &position
}