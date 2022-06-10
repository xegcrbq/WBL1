package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	array := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, elem := range array {
		wg.Add(1)
		go func(elem int) {
			io.WriteString(os.Stdout,
				strconv.Itoa(elem)+"^2="+strconv.Itoa(elem*elem)+" ")
			wg.Done()
		}(elem)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, "\n")
}
