package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	args := os.Args
	if len(args) < 2 {
		log.Fatal("args must more 2")
	}

	port := args[1]
	if len(port) < 2 {
		log.Fatal("port must more 2 length ")
	}

	udpAddr,err:= net.ResolveUDPAddr("udp4",":"+port)
	if err != nil {
		log.Fatal("connect fail,err ",err)
	}

	connect,connErr := net.DialUDP("udp4",nil,udpAddr)
	if err != nil {
		log.Fatal("conenct udp has err",connErr)
	}

	defer connect.Close()

	for {
		reader:=bufio.NewReader(os.Stdin)
		data,_ := reader.ReadString('\n')
		fmt.Println("->",data)
		connect.Write([]byte(data))
		if strings.TrimSpace(data) == "STOP" {
			return
		}
		fmt.Println("\n")
		buffer := make([]byte,1024)
		length,addr,readEr:= connect.ReadFromUDP(buffer)
		if readEr != nil {
			fmt.Printf("->read has err:%s",readEr)
		}
		fmt.Printf("%s->%s",addr,string(buffer[:length]))
	}
}
