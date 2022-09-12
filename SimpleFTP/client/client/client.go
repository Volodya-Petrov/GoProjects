package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type FileInfo struct {
	Path  string
	IsDir bool
}

func List(ip, port, path string) ([]FileInfo, error) {
	conn, err := connectWithServer(ip, port)
	result := []FileInfo{}
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		return result, err

	}
	fmt.Printf("Connected to %v:%v\n", ip, port)
	sendReq(conn, "1", path)
	response, _ := bufio.NewReader(conn).ReadString('\n')
	splitResponse := strings.Split(response[:len(response)-1], " ")
	size, _ := strconv.Atoi(splitResponse[0])
	if size < 0 {
		fmt.Println("File doesn't exist on server")
		return result, os.ErrNotExist
	}
	for i := 1; i <= size*2; i += 2 {
		pathOnServer := splitResponse[i]
		isDir, _ := strconv.ParseBool(splitResponse[i+1])
		result = append(result, FileInfo{pathOnServer, isDir})
	}
	return result, nil
}

func Get(ip, port, pathServer, pathLocal string) (int, error) {
	conn, err := connectWithServer(ip, port)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		return 0, err

	}
	fmt.Printf("Connected to %v:%v\n", ip, port)
	sendReq(conn, "2", pathServer)
	buff := bufio.NewReader(conn)
	sizeInStr, _ := buff.ReadString(' ')
	size, _ := strconv.Atoi(sizeInStr)
	if size < 0 {
		fmt.Println("File doesn't exist on server")
		return 0, os.ErrNotExist
	}
	file, err := os.Create(pathLocal)
	io.Copy(file, conn)
	return size, nil
}

func sendReq(conn net.Conn, command, path string) {
	fmt.Fprintf(conn, "%v %v\n", command, path)
}

func connectWithServer(ip, port string) (net.Conn, error) {
	ipWithPort := net.JoinHostPort(ip, port)
	return net.Dial("tcp", ipWithPort)
}
