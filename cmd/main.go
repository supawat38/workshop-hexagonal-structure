package main

import (
	"fmt"
	protocol "microserviceMOCK/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		fmt.Println(err)
	}
}
