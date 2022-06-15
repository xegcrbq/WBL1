package main

import (
	"fmt"
	"sync"
)

//11.	Что выведет данная программа и почему?
func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, j int) {
			fmt.Println(j)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
