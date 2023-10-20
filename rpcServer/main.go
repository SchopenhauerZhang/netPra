package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

type Data interface{
	Notify(string)
}

type DataRPC struct{
	num int
}

func (data *DataRPC) Notify(s string)  {
	fmt.Println("get person",s)
}

func main(){
	args := os.Args
	if len(args) < 2 {
		log.Fatal("args is less 2")
	}

	if len(args[1]) < 2 {
		log.Fatal("port is nil",args[1])
	}

	data:= new(DataRPC)
	rpc.Register(data)
	fd,err := net.ResolveTCPAddr("tcp4",":"+args[1])
	if err != nil {
		log.Fatal("bind port has err",err)
	}

	l,lErr:= net.ListenTCP("tcp4",fd)
	if lErr != nil {
		log.Fatal("listen err",lErr)
	}

	defer l.Close()

	for {
		c,acpErr:= l.Accept()
		if acpErr != nil {
			log.Print("accept has ",acpErr)
		}

		fmt.Println("rpc connect addr",c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
