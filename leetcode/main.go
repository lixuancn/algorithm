package main

import (
	"fmt"
	"time"
)

func main() {
	run(4)
}

func run(num int) {
	c := make(chan struct{}, 4)
	for i := 0; i < num; i++ {
		c <- struct{}{}
	}
	for {
		select {
		case <-c:
			go job(c)
		}
	}
}

func job(c chan struct{}) {
	defer fmt.Println("job结束")
	fmt.Println("job启动")
	time.Sleep(1 * time.Second)
	c <- struct{}{}
}
