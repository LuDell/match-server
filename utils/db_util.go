package utils

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"github.com/cihub/seelog"
	"fmt"
	"match-server/model"
)

var Conn *gorose.Connection

func init()  {
	Conn = loadCon()
}
func loadCon() *gorose.Connection {
	var config = model.SeeLogConfig.Mysql
	var dbConfig = &gorose.DbConfigSingle{
		Driver:          "mysql",
		EnableQueryLog:  true,
		SetMaxOpenConns: 5,
		SetMaxIdleConns: 3,
		Prefix:          "",
		Dsn:             fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",config.User_name,config.Password,config.Tcp,config.Port,config.Database), // db dsn
	}
	var connection,err = gorose.Open(dbConfig)
	if err != nil {
		seelog.Info("mysql connection error=",err)
	}
	return connection
}