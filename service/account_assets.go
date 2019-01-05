package service

import (
	"github.com/cihub/seelog"
	"match-server/model"
	"match-server/utils"
	"strconv"
)

/**
 *流水类型
 */
const (
	/** 开仓手续费 */
	FEE_TAKER = "fee_taker"
	/** 开仓手续费 */
	FEE_MAKER = "fee_maker"
	/** 平多盈亏 */
	CLOSE_BUY_PROFIT_LOSS = "close_buy_profit_loss"
	/** 平空盈亏 */
	CLOSE_SELL_PROFIT_LOSS = "close_sell_profit_loss"
	/** 撮合成交爆仓单多仓盈利注入风险准备金 */
	INJECT_RISK_BLANCE_BUY = "inject_risk_blance_buy"
	/** 撮合成交爆仓单空仓盈利注入风险准备金 */
	INJECT_RISK_BLANCE_SELL = "inject_risk_blance_sell"
)

func SearchBalance(uid uint, acc_type int,isLock bool) float64 {
	var sql = "select balance from account where uid = ? and type = ?"
	if isLock {
		sql += " from update"
	}

	resultList,_ := utils.DBExchange().SQL(sql,uid,acc_type).QueryString()

	var balance,_ = strconv.ParseFloat(resultList[0]["balance"], 64)
	return balance
}

func UpdateBalance(uid uint, acc_type int, amount float64)  {
	var sql = "update account set balance = balance + ? where uid = ? and type = ?"
	_,err := utils.DBExchange().Exec(sql,amount,uid,acc_type, "xorm")
	if err != nil {
		seelog.Error("account update error, ",err)
	}
	panic(err)
}

func insertTrans(transactions ...*model.Transaction) (int64 ,error) {
	res,err := utils.DBExchange().Insert(transactions)
	if err != nil {
		seelog.Error("transaction insert error",err)
		return 0,nil
	}
	return res,err
}