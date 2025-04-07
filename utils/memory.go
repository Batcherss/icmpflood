package utils

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var memoryLimit int64 = 2 * 1024 * 1024 * 1024 // 2 GB

func MonitorMemoryUsage() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		if int64(m.Alloc) > memoryLimit {
			fmt.Println("memory buffer overflow. exiting...")
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
	}
}
