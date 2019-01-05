package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/server"
	"match-server/service"
)

var (
	addr = flag.String("addr", "127.0.0.1:8997", "server address")
)

type Echo int

func (t *Echo) Echo(ctx context.Context, args []byte, reply *[]byte) error {
	*reply = []byte("hello" + string(args))
	return nil
}

func main() {
	flag.Parse()
	//首次加载账户
	service.LoadGlobalConf()
	s := server.NewServer()
	s.RegisterName("Echo", new(Echo), "")
	s.Serve("tcp", *addr)
}