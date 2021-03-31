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
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 2*time.Second)
	// Пытается установить TCP соединение
	if err == nil {
		println(port)
	}
	// Если удаётся, печатает этот порт
	if port+workersCount < 65536 {
		// В конце запускается на сканирование другого порта, остаток которого по модулю
		// количества горутинов равен номеру горутина.
		scanPort(ip, port+workersCount, workersCount)
	}
}
func scanPortWg(ip string, port int, workersCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	scanPort(ip, port, workersCount)
}

func main() {
	ip := os.Args[1]
	var wg sync.WaitGroup
	for i := 1; i < 2000; i++ { // Делает 2000 горутинов
		wg.Add(1)
		go scanPortWg(ip, i, 2000, &wg) // У каждого горутина запускает сканирование
		//начиная с порта i
	}
	wg.Wait()
}
