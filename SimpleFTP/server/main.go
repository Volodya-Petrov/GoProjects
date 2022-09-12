package main

import (
	"fmt"
)

func main() {
	listener, _ := Run(".", ":1337")
	var exit string
	for exit != "exit" {
		fmt.Scanf("%s", &exit)
	}
	listener.Close()
}
