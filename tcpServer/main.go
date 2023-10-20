package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync/atomic"
	"strings"
	"time"
)

func main()  {
	args:= os.Args
	if len(args) < 1 {
		log.Fatal("must input ")
	}

	listenPort := args[1]
	if len(listenPort) < 2 {
		log.Fatal("port not format",listenPort)
	}

	conn,listenErr:= net.Listen("tcp",":"+listenPort)
	if listenErr != nil {
		log.Fatal("listen port fail,err:",listenErr)
	}
	defer func() {
		if closeErr:=conn.Close();closeErr != nil {
			log.Println("close has err:",closeErr)
		}
	}()

	acpConn,acceptErr := conn.Accept()
	if acceptErr != nil {
		log.Println("accept has err:",acceptErr)
	}

	transferSendBytes := atomic.Uint64{}

	for {
		readContent,readErr := bufio.NewReader(acpConn).ReadString('\n')
		if readErr != nil {
			log.Printf("read has err:",readErr)
		}

		formatContent := strings.TrimSpace(readContent)
		if formatContent == "STOP"{
			log.Printf("client request close connection,port:",listenPort)
			fmt.Println("client close connect")
			return
		}

		fmt.Println("->",string(formatContent))
		sendBytes,sendErr:= acpConn.Write([]byte(fmt.Sprintf("%s send receive time %s \n",listenPort,time.Now().Format(time.RFC3339))))
		if sendErr != nil {
			log.Println("send has err:",sendErr)
		}
		transferSendBytes.Add(uint64(sendBytes))
		fmt.Printf(">> has send %d bytes",transferSendBytes.Load())
	}

}