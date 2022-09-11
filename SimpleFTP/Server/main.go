package main

import (
	server "Server/Server"
	"fmt"
)

func main() {
	listener, _ := server.Run(".", ":1337")
	var exit string
	for exit != "exit" {
		fmt.Scanf("%s", &exit)
	}
	listener.Close()
}
