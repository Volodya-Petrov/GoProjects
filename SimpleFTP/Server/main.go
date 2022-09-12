package main

import (
	"fmt"
	"server/server"
)

func main() {
	listener, err := server.Run(".", "1337")
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	var exit string
	for exit != "exit" {
		fmt.Scanf("%s", &exit)
	}
	listener.Close()
}
