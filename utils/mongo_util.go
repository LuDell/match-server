package utils

import (
	"gopkg.in/mgo.v2"
	"match-server/model"
)

var Session *mgo.Session

func del_init()  {
	Session = globalSession()
}

var config = model.SeeLogConfig.Mongo
func globalSession() *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{config.Tcp+":"+config.Port},
		Direct:    false,
		Timeout:   config.Timeout,
		Username:  config.User_name,
		Password:  config.Password,
		PoolLimit: 5, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	session.SetMode(mgo.Monotonic,true)
	return session
}

func ShowDB() *mgo.Database {
	session := Session.Copy()
	return session.DB(config.Database)
}
