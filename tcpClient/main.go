package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	// 命令行读取输入
	args := os.Args
	if len(args) < 1 {
		log.Fatal("netPra must inpurt args")
		return
	}

	// 命令行读取连接地址
	connectAddr := args[1]
	if len(connectAddr) < 1 {
		log.Fatal("connectAddr not nil")
		return
	}

	// tcp连接
	conn,err:= net.Dial("tcp",connectAddr)
	if err != nil {
		log.Fatal("connect fail,err:",err)
		return
	}

	for {
		// 从控制台读取输入
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("<<")
		// 以\n为分隔符
		content,readErr:=reader.ReadString('\n')
		if readErr != nil {
			log.Println("readErr",readErr)
		}

		fmt.Printf("本地输入：%s",content)
		receiveMsg,receviErr :=bufio.NewReader(conn).ReadString('\n')
		fmt.Println("->:%s",receiveMsg)
		if receviErr != nil {
			log.Fatal("receive Err:",receviErr)
			return
		}
	}
}
