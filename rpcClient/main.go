package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)


type Person struct {
	num int
}

func (person *Person) Notify(string) {

}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("args is less 2")
	}

	fd, diErr := rpc.Dial("tcp4", ":"+args[1])
	if diErr != nil {
		log.Fatal("has err ", diErr)
	}

	callErr := fd.Call("DataRPC.Notify","client come",nil)
	if callErr != nil {
		fmt.Println("-> has err",callErr)
	}
	fmt.Println("-> call success")
}
