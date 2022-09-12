package server

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

func Run(path, port string) (net.Listener, error) {
	file, errorFromStat := os.Stat(path)
	if errorFromStat != nil {
		return nil, errorFromStat
	}
	if !file.IsDir() {
		return nil, os.ErrNotExist
	}
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Слушаем порт:%v\n", port)
	go start(listener)
	return listener, nil
}

func start(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка с подключением:%v\n", err)
			continue
		}
		fmt.Printf("Подключение установлено с %v\n", conn.RemoteAddr())
		go workWithClient(conn)
	}
}

func workWithClient(conn net.Conn) {
	buff := bufio.NewReader(conn)
	readReq, _ := buff.ReadString('\n')
	splitReq := strings.Split(readReq[:len(readReq)-1], " ")
	if splitReq[0] == "1" {
		listReq(conn, splitReq[1])
	}
	if splitReq[0] == "2" {
		getReq(conn, splitReq[1])
	}
	fmt.Printf("Закрываем подключение с %v\n", conn.RemoteAddr())
	conn.Close()
}

func listReq(conn net.Conn, path string) {
	info := &strings.Builder{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(conn, "-1 ")
		return
	}
	for i, file := range files {
		info.WriteString(fmt.Sprintf("%v %v", file.Name(), file.IsDir()))
		if i != len(files)-1 {
			info.WriteString(" ")
		}
	}
	fmt.Fprintf(conn, strconv.Itoa(len(files))+" "+info.String()+"e\n")
}

func getReq(conn net.Conn, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(conn, "-1 ")
		return
	}
	fileInfo, _ := file.Stat()
	fmt.Fprintf(conn, string(fileInfo.Size())+" ")
	io.Copy(conn, file)
}
