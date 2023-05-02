package main

import (
	"fmt"
	"time"
)

func procLoop(dat1 []interface{}, duration int) {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Duration(duration) * time.Millisecond)
		fmt.Println(dat1, i)
	}
}

func main() {
	data1 := []interface{}{
		"Coba1", "Coba2", "Coba3",
	}
	data2 := []interface{}{
		"Bisa1", "Bisa2", "Bisa3",
	}

	go procLoop(data1, 444)
	go procLoop(data2, 777)
	time.Sleep(3000 * time.Millisecond)
}
