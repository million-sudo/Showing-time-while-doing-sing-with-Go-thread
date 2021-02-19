package main

import (
	"fmt"
	"time"
)

func sing(t1 chan bool) {
	for i := 0;i < 5;i++ {
		fmt.Println("La la la")
		time.Sleep(100 * time.Millisecond)
	}
	t1 <- true
	
}

func show_time(t2 chan bool) {
	for i := 0;i < 5;i++ {
		fmt.Println(time.Now())
		time.Sleep(100 * time.Millisecond)
	}
	t2 <- true
}

func main() {
	var thread1 chan bool = make(chan bool)
	var thread2 chan bool = make(chan bool)

	// start sing() thread
	go sing(thread1)
	// start show_time() thread
	go show_time(thread2)
	
	// wait the routins
	s := <- thread1
	t := <- thread2
	fmt.Println(s,t)
}
