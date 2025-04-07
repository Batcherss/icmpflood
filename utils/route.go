package utils

import (
	"fmt"
	"sync/atomic"
	"time"
	"sync"
)

var udpCount int64
var icmpCount int64
var totalBytes int64
var routeStress int64

var mu sync.Mutex  

func InfoBar() {
	for {
		time.Sleep(time.Second)

		bytes := atomic.LoadInt64(&totalBytes)
		stress := atomic.LoadInt64(&routeStress)

		kb := bytes / 1024

		mu.Lock()
		fmt.Printf("send kb's: %d KB\n", kb)
		fmt.Printf("total bytes sent: %d bytes\n", bytes)
		fmt.Printf("ping: %d ms (%.2f%%)\n", stress, float64(stress)/50)
		mu.Unlock()
	}
}
