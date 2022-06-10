package main

import (
	"context"
	"fmt"
	"time"
)

/* Реализовать все возможные способы остановки выполнения горутины.  */

func worker1(ch chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("Горутина 1 стоп")
			return
		default:
			fmt.Println("Горутина 1 работает")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func worker2(ch chan int) {
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("Горутина 2 стоп")
			return
		}
		fmt.Println("Горутина 2 работает ")
	}
}

func worker3(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 3 стоп")
				close(ch)
				return
			case ch <- n:
				time.Sleep(time.Second)
				//fmt.Println("Горутина 3 работает:", n)
				n++
			}
		}
	}()
	return ch
}
func main() {

	// 1 способ горутина работает, пока не получит данные из канала
	ch1 := make(chan int)
	go worker1(ch1)
	time.Sleep(1 * time.Second)
	ch1 <- 1

	// 2 способ горутина работает, пока канал не закрыт
	ch2 := make(chan int)
	go worker2(ch2)
	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)
	//time.Sleep(5 * time.Second)
	//go worker2(ch2)

	// 2 способ блокировка горутины с помощью контекста Deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	time.Sleep(4 * time.Second)
	for n := range worker3(ctx) {
		fmt.Println("Горутина 3 работает: ", n)
	}
}
