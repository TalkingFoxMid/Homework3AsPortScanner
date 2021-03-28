package main

import (
	"fmt"
	_ "fmt"
	_ "log"
	"net"
	_ "net"
	"os"
	_ "os"
	_ "strconv"
	_ "strings"
	"sync"
	"time"
)

func scanPort(ip string, port int, workersCount int) {

	_,  err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 2*time.Second)
	if (err == nil) {
		println(port)
	}
	if port + workersCount < 65536 {
		scanPort(ip, port + workersCount, workersCount)
	}
}
func scanPortWg(ip string, port int, workersCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	scanPort(ip, port, workersCount)
}

func main() {
	ip := os.Args[1]
	var wg sync.WaitGroup
	for i := 1; i < 2000; i++ {
		wg.Add(1)
		go scanPortWg(ip, i, 2000, &wg)
	}
	wg.Wait()
}