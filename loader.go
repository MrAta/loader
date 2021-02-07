package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var logger *log.Logger
var addr string

func main() {
	addr = os.Getenv("HTTP_ADDR")
	fileNamePtr := flag.String("logfile", "info.log", "Log file name")
	ratePtr := flag.Int("rate", 1, "Request rate per second")
	durationPtr := flag.String("duration", "1s", "Test duration: Ns, Nm, Nh")
	flag.Parse()

	duration := getDuration(*durationPtr)

	file, err := os.OpenFile(*fileNamePtr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Ltime)
	fmt.Println("Start: ", time.Now())
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	finishedLaunch := make(chan bool)

	go launchWorkers(ratePtr, done, finishedLaunch, ticker, logger)
	time.Sleep(duration)
	ticker.Stop()
	done <- true

	<-finishedLaunch
	fmt.Println("End: ", time.Now())
}

func launchWorkers(ratePtr *int, done chan bool, finishedLaunch chan bool, ticker *time.Ticker, logger *log.Logger) {
	var wg sync.WaitGroup
	for {
		select {
		case <-done:
			goto WaitForWorkers
		case <-ticker.C:
			wg.Add(1)
			go sendReq(addr, *ratePtr, &wg, time.Now(), logger)
		}
	}
WaitForWorkers:
	wg.Wait()
	finishedLaunch <- true
}
func sendReq(addr string, rate int, wg *sync.WaitGroup, t time.Time, logger *log.Logger) {

	var reqWg sync.WaitGroup
	for i := 0; i < rate; i++ {
		reqWg.Add(1)
		go func(addr string, l *log.Logger, t time.Time, reqWg *sync.WaitGroup) {
			defer reqWg.Done()
			resp, err := http.Get(addr)
			var statusCode int
			if err != nil {
				statusCode = 500
			} else {
				statusCode = resp.StatusCode
			}

			l.Println(t, statusCode)
			//fmt.Println(t, resp.StatusCode)
		}(addr, logger, t, &reqWg)
	}
	reqWg.Wait()
	wg.Done()
}

func getDuration(d string) time.Duration {
	duration, e := strconv.Atoi(d[:len(d)-1])
	if e != nil {
		duration = 1
	}
	switch d[len(d)-1] {
	case 's':
		return time.Duration(duration) * time.Second
	case 'm':
		return time.Duration(duration) * time.Minute
	case 'h':
		return time.Duration(duration) * time.Hour
	default:
		return time.Second
	}
}
