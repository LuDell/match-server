package utils

import (
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func init() {
	defer seelog.Flush()
	//加载配置文件
	Logger, _ = seelog.LoggerFromConfigAsFile("config/seelog.xml")

	//替换记录器
	seelog.ReplaceLogger(Logger)
}
