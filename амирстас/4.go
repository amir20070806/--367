package main

import (
 "fmt"
 "math/rand"
 "sync"
 "time"
)

func main() {
 candidates := map[string]int{
  "Кандидат Амир": 0,
  "Кандидат Стас": 0,
  "Кандидат Владимир Путин": 0,
 }
 var mu sync.Mutex
 
 names := []string{"Кандидат Амир", "Кандидат Стас", "Кандидат Владимир Путин"}
 
 fmt.Println("Начало голосования!")

 for i := 0; i < 3; i++ {
  go func(id int) {
   for j := 0; j < 5; j++ {
    time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
    
    candidate := names[rand.Intn(len(names))]
    
    mu.Lock()
    candidates[candidate]++
    fmt.Printf("Горутина %d: голос за %s\n", id, candidate)
    mu.Unlock()
   }
  }(i)
 }

 time.Sleep(2 * time.Second)
 
 mu.Lock()
 fmt.Println("результэт")
 for candidate, votes := range candidates {
  fmt.Printf("%s: %d голосов\n", candidate, votes)
 }
 mu.Unlock()
}