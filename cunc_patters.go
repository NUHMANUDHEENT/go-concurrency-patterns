package main

import (
	"fmt"
	"sync"
)

func PrintNumbers(n int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- n
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 1; i < 8; i++ {
		wg.Add(1)
		go PrintNumbers(i, &wg, ch)
	}
	
	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()
	close(ch)
}
