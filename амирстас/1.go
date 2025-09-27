package main 
import ("sync"
"fmt"
)
	var mu sync.Mutex
	var count int
func bray() {
	mu.Lock()
	count = count + 1
	mu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	for i:=0; i < 50; i++ {
		wg.Add(1)
	go func() {
		defer wg.Done()
		bray()
	}()
	}
	wg.Wait()
	fmt.Println("жДОМ:", count)
}