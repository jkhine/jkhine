package main

import (
	"net"
	"strings"
	"strconv"
	"log"
	"time"
	"fmt"
)

func main() {
	addr := strings.Join([]string{"wwww.google.com", strconv.Itoa(20)}, ":")
	timeOut := time.Duration(3) * time.Second
	conn, err := net.DialTimeout("tcp", addr, timeOut)

	if err != nil {
		log.Fatalln(err)
		return
	} else {
		fmt.Println("connection successful")
	}
	defer conn.Close()
}