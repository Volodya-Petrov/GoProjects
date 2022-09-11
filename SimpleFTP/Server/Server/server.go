package Server

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
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
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	go Start(listener)
	return listener, nil
}

func Start(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go WorkWithClient(conn)
	}
}

func WorkWithClient(conn net.Conn) {
	buff := bufio.NewReader(conn)
	readReq, _ := buff.ReadString('\n')
	splitReq := strings.Split(readReq, " ")
	if splitReq[0] == "1" {
		ListReq(conn, splitReq[1])
	}
	if splitReq[0] == "2" {

	}
	conn.Close()
}

func ListReq(conn net.Conn, path string) {
	info := &strings.Builder{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(conn, "-1 ")
		return
	}
	for i, file := range files {
		info.WriteString(fmt.Sprintf("%v, %v", file.Name(), file.IsDir()))
		if i != len(files)-1 {
			info.WriteString(" ")
		}
	}
	fmt.Fprintf(conn, string(len(files))+" "+info.String()+"\n")
}

func GetReq(conn net.Conn, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(conn, "-1 ")
		return
	}
	fileInfo, _ := file.Stat()
	fmt.Fprintf(conn, string(fileInfo.Size())+" ")
	io.Copy(conn, file)
}
