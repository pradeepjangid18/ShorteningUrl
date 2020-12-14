package main

import "sync"

type globalCounter struct {
	mu sync.Mutex
	cnt int64
}

var counter *globalCounter

func(cnt *globalCounter) incrementAndGet() int64{
	cnt.mu.Lock()
	cnt.mu.Unlock()

	cnt.cnt +=1
	return cnt.cnt
}
