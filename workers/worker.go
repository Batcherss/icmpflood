package workers

import (
	"rtstress/handlers"
	"sync"
	"sync/atomic"
	"time"
)

var activeWorkers int32
var maxWorkers int32 = 50

func StartWorker(ip string, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt32(&activeWorkers, 1)
	defer atomic.AddInt32(&activeWorkers, -1)

	if atomic.LoadInt32(&activeWorkers) > maxWorkers {
		time.Sleep(10 * time.Millisecond)
	}

	for {
		select {
		case <-time.After(10 * time.Millisecond): 
			handlers.SendUDPRequest(ip)
			handlers.SendICMPRequest(ip)
		}
	}
}
