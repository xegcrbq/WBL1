package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func read2(id int, ch chan int) { // функция отслеживания воркеров
	for {
		value := <-ch
		fmt.Printf("readed: '%v'\n", value)
	}
}
func main() {
	N := 3
	ch := make(chan int)
	go read2(1, ch)

	timeout := time.After(time.Duration(N) * time.Second) //создание таймера для заверщения программы

	func() {
		var i int
		for {
			i++
			select {
			case <-timeout:
				close(ch)
				return
			default:
				time.Sleep(time.Second)
				ch <- i
			}
		}
	}()
	fmt.Print("stop")
}
