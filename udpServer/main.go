package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	args := os.Args
	if len(args) < 2 {
		log.Fatal("args must more 1 ")
	}

	port := args[1]
	if len(port) < 1 {
		log.Printf("port:%s is not match ",port)
	}

	udpConn,err:= net.ResolveUDPAddr("udp4",":"+port)
	if err != nil {
		log.Fatal("udp connect err:",err)
	}

	l,errListen := net.ListenUDP("udp4",udpConn)
	if errListen != nil {
		log.Fatal("listen udp has err",errListen)
	}
	defer l.Close()

	readBuffer := make([]byte,1024)
	for {
		length,clientAddr,readErr:= l.ReadFromUDP(readBuffer)
		if readErr != nil {
			log.Println("read from udp has err ",readErr)
		}

		fmt.Println("->",string(readBuffer[0:length]))
		fmt.Println("resolve byte:",length)

		l.WriteToUDP([]byte("get it"),clientAddr)
	}

}
