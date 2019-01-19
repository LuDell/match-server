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

	Sell = "SELL"
	Buy = "BUY"
)

func SearchBalance(uid uint, accType int,isLock bool) (account float64,err error) {
	var sql = "select balance from account where uid = ? and type = ?"
	if isLock {
		sql += " from update"
	}

	resultList,err := utils.DBExchange().SQL(sql,uid, accType).QueryString()
	if err != nil || resultList == nil {
		seelog.Error("account is not find,uid=",uid," acc_type= ", accType)
		return 0,errors.New("account is not find")
	}
	balance,err := strconv.ParseFloat(resultList[0]["balance"], 64)
	return balance,err
}

func UpdateBalance(uid uint, accType int, amount float64) error {
	var sql = "update account set balance = balance + ? where uid = ? and type = ?"
	_,err := utils.DBExchange().Exec(sql,amount, uid, accType)
	if err != nil {
		seelog.Error("account update error, ",err)
	}
	return err
}

func insertTrans(transactions ...*model.Transaction) error {
	//一笔一笔插入流水
	for _,tran := range transactions {
		var fromBalance,toBalance float64
		var err1,err2 error
		//顺序加锁
		if tran.ToType > tran.FromType {
			toBalance,err2 = SearchBalance(tran.ToUid,tran.ToType,true)
			fromBalance,err1 = SearchBalance(tran.FromUid,tran.FromType,true)
		}else{
			fromBalance,err1 = SearchBalance(tran.FromUid,tran.FromType,true)
			toBalance,err2 = SearchBalance(tran.ToUid,tran.ToType,true)
		}
		if err1 == nil{
			seelog.Info("用户",tran.FromUid,"的", tran.FromType, "账户异常")
			return err1
		}
		if err2 == nil{
			seelog.Info("用户",tran.ToUid,"的", tran.ToType, "账户异常")
			return err2
		}
		tran.FromBalance = fromBalance - tran.Amount
		tran.ToBalance = toBalance + tran.Amount

		//1.插入数据
		_,err := utils.DBExchange().Insert(tran)
		if err != nil {
			seelog.Error("transaction insert error",err)
			return err
		}
		//2.更新账户
		if err := UpdateBalance(tran.FromUid, tran.FromType, -tran.Amount); err!= nil {
			return err
		}
		if err := UpdateBalance(tran.ToUid, tran.ToType, tran.Amount); err!= nil {
			return err
		}
	}
	return nil
}