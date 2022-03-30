package main

import (
	"context"
	"sync"
	"time"
)

var wg sync.WaitGroup{}
func main() {
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	wg.Add(2)
	go B(ctx)
	select
	case <-chan
	go C()
}

func B() bool{

	defer wg.Done()

	chan <- true

	return true
}

func C(ctx context.Context){

	time.Sleep(10 * time.Second)
}