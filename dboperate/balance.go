package dboperate

import (
	"match-server/utils"
	"strconv"
)

const baseSQL = ""

func SearchInfo(uid uint, acc_type uint,isLock bool) float64 {
	var sql = "select balance from account where uid = ? and type = ?"
	if isLock {
		sql += " from update"
	}

	resultList,_ := utils.DBExchange().SQL(sql,uid,acc_type).QueryString()

	var balance,_ = strconv.ParseFloat(resultList[0]["balance"], 64)
	return balance
}
