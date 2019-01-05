package service

import (
	"github.com/cihub/seelog"
	"github.com/pkg/errors"
	"match-server/model"
	"match-server/utils"
	"strconv"
)

/**
 *流水类型
 */
const (
	/** 开仓手续费 */
	FeeTaker = "fee_taker"
	/** 开仓手续费 */
	FeeMaker = "fee_maker"
	/** 平多盈亏 */
	CloseBuyProfitLoss = "close_buy_profit_loss"
	/** 平空盈亏 */
	CloseSellProfitLoss = "close_sell_profit_loss"
	/** 撮合成交爆仓单多仓盈利注入风险准备金 */
	InjectRiskBalanceBuy = "inject_risk_balance_buy"
	/** 撮合成交爆仓单空仓盈利注入风险准备金 */
	InjectRiskBalanceSell = "inject_risk_balance_sell"
)

func SearchBalance(uid uint, acc_type int,isLock bool) (account float64,err error) {
	var sql = "select balance from account where uid = ? and type = ?"
	if isLock {
		sql += " from update"
	}

	resultList,err := utils.DBExchange().SQL(sql,uid,acc_type).QueryString()
	if err != nil || resultList == nil {
		seelog.Error("account is not find,uid=",uid," acc_type= ",acc_type)
		return 0,errors.New("account is not find")
	}
	balance,err := strconv.ParseFloat(resultList[0]["balance"], 64)
	return balance,err
}

func UpdateBalance(uid uint, acc_type int, amount float64) error {
	var sql = "update account set balance = balance + ? where uid = ? and type = ?"
	_,err := utils.DBExchange().Exec(sql,amount,uid,acc_type, "xorm")
	if err != nil {
		seelog.Error("account update error, ",err)
	}
	return err
}

func insertTrans(transactions ...*model.Transaction) error {
	_,err := utils.DBExchange().Insert(transactions)
	if err != nil {
		seelog.Error("transaction insert error",err)
	}
	for _,tran := range transactions {
		if err := UpdateBalance(tran.FromUid, tran.FromType, -tran.Amount); err!= nil {
			return err
		}
		if err := UpdateBalance(tran.ToUid, tran.ToType, tran.Amount); err!= nil {
			return err
		}
	}
	return err
}