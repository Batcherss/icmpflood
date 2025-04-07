package handlers

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"math/rand"
	"net"
	"os"
)

var icmpCount int64

func SendICMPRequest(ip string) {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("err with connect icmp:", err)
		return
	}
	defer c.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq: rand.Intn(65535),
			Data: make([]byte, 1024),
		},
	}

	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		fmt.Println("err with encode icmp:", err)
		return
	}

	_, err = c.WriteTo(msgBytes, &net.IPAddr{IP: net.ParseIP(ip)})
	if err != nil {
		fmt.Println("err with send icmp:", err)
		return
	}

	fmt.Printf("[+] Send ICMP\n")
}
