package utils

import (
	"fmt"
	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func DBExchange() *xorm.Engine {
	var config = seeLogConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",config.User_name,config.Password,config.Tcp,config.Port,"exchange")

	engine, err := xorm.NewEngine("mysql", dsn)
	if(err != nil){
		seelog.Error("database connection is error",err)
	}
	engine.ShowSQL(config.ShowLog)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	return engine
}

func DBContract() *xorm.Engine {
	var config = seeLogConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",config.User_name,config.Password,config.Tcp,config.Port,"contract")

	engine, err := xorm.NewEngine("mysql", dsn)
	if(err != nil){
		seelog.Error("database connection is error",err)
	}
	engine.ShowSQL(config.ShowLog)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	return engine
}