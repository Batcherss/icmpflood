package main

import (
	"fmt"
	"sync"
	"rtstress/utils"
	"rtstress/workers"
)

func main() {
	var ip string
	fmt.Print("ip: ")
	fmt.Scanln(&ip)

	go utils.InfoBar()
	go utils.MonitorMemoryUsage()

	var wg sync.WaitGroup

	numThreads := 50

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go workers.StartWorker(ip, &wg)
	}

	select {}
}
