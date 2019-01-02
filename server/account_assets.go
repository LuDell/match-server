package server

import (
	"github.com/cihub/seelog"
	"match-server/utils"
	"strconv"
)


func SearchInfo(uid uint, acc_type uint,isLock bool) float64 {
	var sql = "select balance from account where uid = ? and type = ?"
	if isLock {
		sql += " from update"
	}

	resultList,_ := utils.DBExchange().SQL(sql,uid,acc_type).QueryString()

	var balance,_ = strconv.ParseFloat(resultList[0]["balance"], 64)
	return balance
}

func UpdateBalance(uid uint, acc_type uint, amount float64)  {
	var sql = "update account set balance = balance + ? where uid = ? and type = ?"
	_,err := utils.DBExchange().Exec(sql,amount,uid,acc_type, "xorm")
	if err != nil {
		seelog.Error("account update error, ",err)
	}
	panic(err)
}