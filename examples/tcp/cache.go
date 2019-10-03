package main

import "sync"

var mu sync.RWMutex

type cache map[string]string

func (c cache) Set(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	c[key] = value
}
func (c cache) Get(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return c[key]
}
