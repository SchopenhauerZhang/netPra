package main 

import "fmt"

import "net"

func main() {
  fmt.println("this is websocket")
}

func server() {
   fmt.println("this is websocket server ")
}

func run() {
  var con *net.Conn
  con:= net.Connect("tcp","9090")
  con.listen()
  con.bind()
  con.accept()
}



