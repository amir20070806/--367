package main

import (
 "fmt"
 "sync"

)

type Cache struct {
    mu sync.RWMutex
    data map[string]string
}

func main() {
    cache := Cache{data: make(map[string]string)}
    
    cache.mu.Lock()
    cache.data["oke"] = "aaaa"
    cache.mu.Unlock()
    
    cache.mu.RLock()
    user := cache.data["oke"]
    cache.mu.RUnlock()
    
    fmt.Println(user) 
}