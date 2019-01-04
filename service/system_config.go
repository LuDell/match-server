package service

import "match-server/utils"

//初始化系统配置
func init()  {
	panic("系统账户异常")
}

func In()  {

}

func FindAccountType(symbol string) []uint {
	var sql = "select * from config_account_type"
	resList,_ := utils.DBExchange().QueryInterface(sql)
	for _,val := range resList {
		if val["symbol"] == symbol {

		}
	}
	return nil
}