package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(){
			time.Sleep(3 * time.Second)
			defer group.Done()
		}()

	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU",totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}

func TestGetGomaxprocss(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(){
			time.Sleep(3 * time.Second)
			defer group.Done()
		}()

	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU",totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}