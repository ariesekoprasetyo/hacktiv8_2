package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func procLoop1(dat1 []interface{}, wg *sync.WaitGroup) {
	for i := 1; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(dat1, i)

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

	mutex.Lock()
	go procLoop1(data1, &wg)
	mutex.Unlock()

	mutex.Lock()
	go procLoop1(data2, &wg)
	mutex.Unlock()

	wg.Wait()
}
