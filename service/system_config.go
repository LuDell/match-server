package service

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"match-server/utils"
	"strconv"
	"strings"
)

const (
	CExchangeFee   = "108"
	CRiskAssure    = "109"
	CPositionClose = "111"
	UMargin        = "216"
)

//系统交易账户
var accountType = map[string]map[string]int{}
//系统用户
var sysUid = uint(1084)

func LoadGlobalConf() {
	var err error
	var sql = "select * from config_account_type"
	resList,err := utils.DBExchange().QueryString(sql)
	if err != nil {
		seelog.Error("账户获取失败",err)
		panic("load account type is error")
	}
	//手续费账户
	var cExchangeFee = map[string]int{}
	//风险准备金账户
	var cRiskAssure = map[string]int{}
	//平仓账户
	var cPositionClose = map[string]int{}
	//保证金账户
	var uMargin = map[string]int{}
	for _,val := range resList {
		var assetType,_ = strconv.Atoi(val["asset_type"])
		if strings.HasPrefix(val["asset_type"], UMargin) {
			uMargin[val["coin_symbol"]] = assetType
		}
		if strings.HasPrefix(val["asset_type"], CExchangeFee) {
			cExchangeFee[val["symbol"]] = assetType
		}
		if strings.HasPrefix(val["asset_type"], CRiskAssure) {
			cRiskAssure[val["coin_symbol"]] = assetType
		}
		if strings.HasPrefix(val["asset_type"], CPositionClose) {
			cPositionClose[val["symbol"]] = assetType
		}
	}
	accountType[UMargin] = uMargin
	accountType[CExchangeFee] = cExchangeFee
	accountType[CRiskAssure] = cRiskAssure
	accountType[CPositionClose] = cPositionClose
	jsonByte,_ := json.Marshal(accountType)
	seelog.Info("load account type ",string(jsonByte))
}