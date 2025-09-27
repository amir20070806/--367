package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {
 var mu sync.Mutex
 var wg sync.WaitGroup
 
 log := func(msg string) {
  mu.Lock()
  defer mu.Unlock()
  fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), msg)
 }
 
 // Запускаем горутины
 for i := 1; i <= 3; i++ {
  wg.Add(1)
  go func(id int) {
   defer wg.Done()
   for j := 1; j <= 3; j++ {
    log(fmt.Sprintf("Поток %d - Сообщение %d", id, j))
    time.Sleep(50 * time.Millisecond)
   }
  }(i)
 }
 
 wg.Wait()
 log("Программа завершена")
}
