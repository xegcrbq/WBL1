package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	array := make([]int, 1000)
	for i := range array {
		array[i] = 2
	}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	sum := 0
	for _, elem := range array {
		wg.Add(1)
		go func(elem int) {
			m.Lock()
			sum += elem * elem
			m.Unlock()
			wg.Done()
		}(elem)
	}
	wg.Wait()
	fmt.Println(sum)
}
