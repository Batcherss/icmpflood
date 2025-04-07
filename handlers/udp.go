package handlers

import (
	"fmt"
	"net"
)

var udpCount int64
var totalBytes int64

func SendUDPRequest(ip string) {
	conn, err := net.Dial("udp", ip+":12345")
	if err != nil {
		fmt.Println("err with stable udp protocol:", err)
		return
	}
	defer conn.Close()

	message := make([]byte, 1024) 
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("err with send udp:", err)
		return
	}

	fmt.Printf("[+] Send UDP\n")
}
