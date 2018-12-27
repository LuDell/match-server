package utils

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	closeChan chan byte
	once sync.Once
}

func InitConnection(wsConn *websocket.Conn)(conn *Connection) {
	conn = &Connection{
		wsConn: wsConn,
		inChan: make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),
		closeChan: make(chan byte),
		once: sync.Once{},
	}

	go conn.ReadLoop()
	go conn.WriteLoop()

	return conn
}

func (conn *Connection)ReadMessage()(data []byte,err error)  {
	select {
	case data = <- conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return data,err
}

func (conn *Connection)WriteMessage(data []byte)(err error)  {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return err
}


func (conn *Connection)Close()  {
	conn.wsConn .Close()

	//代码只执行一次
	conn.once.Do(func() {
		seelog.Info("connection is closed")
		close(conn.closeChan)
	})
}

func (conn *Connection) ReadLoop() {
	var(
		data []byte
		err error
	)
	for{
		if _,data,err = conn.wsConn.ReadMessage(); err != nil{
			goto Err
		}
		select {
		case conn.inChan <- data:
			seelog.Info("write message>>>",string(data))
		case <- conn.closeChan:
			goto Err
		}

	}
	Err:
		conn.Close()
}

func (conn *Connection) WriteLoop() {
	var(
		data []byte
		err error
	)
	for{
		select {
		case data = <- conn.outChan:
			seelog.Info("read message>>>",string(data))
		case <- conn.closeChan:
			goto Err
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage,data); err != nil{
			goto Err
		}
	}
Err:
	conn.Close()
}