package service

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"match-server/utils"
	"strconv"
	"strings"
)

const (
	C_EXCHANGE_FEE = "108"
	C_RISK_ASSURE = "109"
	C_POSITION_CLOSE = "111"
	U_MARGIN = "216"
)
var accountType = map[string]map[string]int{}

func LoadAccountType() {
	var err error
	var sql = "select * from config_account_type"
	resList,err := utils.DBExchange().QueryString(sql)
	if(err != nil){
		seelog.Error("账户获取失败",err)
		panic("load account type is error")
	}
	//手续费账户
	var c_exchange_fee = map[string]int{}
	//风险准备金账户
	var c_risk_assure = map[string]int{}
	//平仓账户
	var c_position_close = map[string]int{}
	//保证金账户
	var u_margin = map[string]int{}
	for _,val := range resList {
		var asset_type,_ = strconv.Atoi(val["asset_type"]);
		if strings.HasPrefix(val["asset_type"],U_MARGIN) {
			u_margin[val["coin_symbol"]] = asset_type
		}
		if strings.HasPrefix(val["asset_type"],C_EXCHANGE_FEE) {
			c_exchange_fee[val["symbol"]] = asset_type
		}
		if strings.HasPrefix(val["asset_type"],C_RISK_ASSURE) {
			c_risk_assure[val["coin_symbol"]] = asset_type
		}
		if strings.HasPrefix(val["asset_type"],C_POSITION_CLOSE) {
			c_position_close[val["symbol"]] = asset_type
		}
	}
	accountType[U_MARGIN] = u_margin
	accountType[C_EXCHANGE_FEE] = c_exchange_fee
	accountType[C_RISK_ASSURE] = c_risk_assure
	accountType[C_POSITION_CLOSE] = c_position_close
	jsonByte,_ := json.Marshal(accountType);
	seelog.Info("load account type ",string(jsonByte))
}