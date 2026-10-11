package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const DEFAULT_WORKER_CAP = 2

func inbound(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int, 256)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()

	return out
}

func sq_Nchan_transformer(ctx context.Context, in <-chan int, worker int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 256)

	w := DEFAULT_WORKER_CAP
	if worker > DEFAULT_WORKER_CAP {
		w = worker
	}

	for i := 0; i < w; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range in {
				select {
				case <-ctx.Done():
					return
				case out <- val * val:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	cap := 10000000
	nums := make([]int, cap+1)

	ctx, cancle := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancle()

	for i := 0; i <= cap; i++ {
		nums[i] = i
	}

	for n := range sq_Nchan_transformer(ctx, inbound(ctx, nums...), 6) {
		fmt.Println(n)
	}

	fmt.Println("hello world")
}
