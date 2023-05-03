package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var flag bool

func procLoop1(dat1 []interface{}, isEven bool) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		if flag == isEven {
			fmt.Println(dat1, i)
			flag = !isEven
		}
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	data1 := []interface{}{
		"Coba1", "Coba2", "Coba3",
	}
	data2 := []interface{}{
		"Bisa1", "Bisa2", "Bisa3",
	}
	wg.Add(2)
	flag = false
	go procLoop1(data1, true)
	go procLoop1(data2, false)
	wg.Wait()
}
