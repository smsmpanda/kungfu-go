package network

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func Gorun2() {
	calc := new(Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
