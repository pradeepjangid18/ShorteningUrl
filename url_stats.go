package main

import "sync"

type urlStats struct {
	mu sync.Mutex
	urlToCount map[string]int
}

var stats *urlStats

func(st *urlStats) updateUrlStats(url string) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.urlToCount[url]+=1
}

func(st *urlStats) getUrlHitCount(url string) int {
	st.mu.Lock()
	defer st.mu.Unlock()

	return st.urlToCount[url]
}
