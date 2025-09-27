package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {

 queue := []string{}
 var mu sync.Mutex
 

 go func() {
  mu.Lock()
  queue = append(queue, "Задача адын")
  fmt.Println("Добавлена задача адын")
  mu.Unlock()
 }()
 
 go func() {
  mu.Lock()
  queue = append(queue, "Задача двэ") 
  fmt.Println("Добавлена задача двэ")
  mu.Unlock()
 }()
 

 go func() {
  mu.Lock()
  if len(queue) > 0 {
   task := queue[0]
   queue = queue[1:]
   fmt.Println("Обработана:", task)
  }
  mu.Unlock()
 }()
 

 time.Sleep(100 * time.Millisecond)
 
 mu.Lock()
 fmt.Println("Осталось задач:", len(queue))
 mu.Unlock()
}