package main

import (
 "fmt"
 "time"
)

func main() {
 fmt.Println("Кэш БЕЗ обновления TTL при чтении    ")
 
 // Простой кэш без обновлений
 cache1 := make(map[string]time.Time)
 cache1["data"] = time.Now().Add(2 * time.Second)
 
 for i := 0; i < 4; i++ {
  if expiry, exists := cache1["data"]; exists {
   if time.Now().Before(expiry) {
    fmt.Printf("Данные существуют (TTL не обновляется)    ")
   } else {
    fmt.Printf("Данные истекли    ")
    delete(cache1, "data")
   }
  }
  time.Sleep(1 * time.Second)
 }
 
 fmt.Println(" Кэш С обновлением TTL при чтении    ")
 
 // Кэш с обновлением
 cache2 := make(map[string]time.Time)
 ttl := 2 * time.Second
 cache2["data"] = time.Now().Add(ttl)
 
 for i := 0; i < 6; i++ {
  if expiry, exists := cache2["data"]; exists {
   if time.Now().Before(expiry) {
    fmt.Printf("Данные существуют (TTL обновлен)    ")
    cache2["data"] = time.Now().Add(ttl)
   } else {
	
    fmt.Printf("Данные истекли    ")
    delete(cache2, "data")
   }
  }
  time.Sleep(1 * time.Second)
 }
}