package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {
 iphoneStock := 10
 var mu sync.Mutex
 

 go func() {
  for i := 0; i < 3; i++ {
   mu.Lock()
   iphoneStock += 5
   fmt.Printf("Поставка: iPhone +5 = %d\n", iphoneStock)
   mu.Unlock()
   time.Sleep(200 * time.Millisecond)
  }
 }()
 

 for i := 0; i < 3; i++ {
  go func(id int) {
   for j := 0; j < 2; j++ {
    mu.Lock()
    if iphoneStock >= 2 {
     iphoneStock -= 2
     fmt.Printf("Продажа %d: iPhone -2 = %d\n", id, iphoneStock)
    } else {
     fmt.Printf("Продажа %d: недостаточно iPhone\n", id)
    }
    mu.Unlock()
    time.Sleep(150 * time.Millisecond)
   }
  }(i)
 }
 
 time.Sleep(1 * time.Second)
 mu.Lock()
 fmt.Printf("Итоговый остаток: %d iPhone\n", iphoneStock)
 mu.Unlock()
}