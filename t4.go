package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func read(id int, ch chan int) { // функция отслеживания воркеров
	for {
		value, ok := <-ch
		if !ok { //проверка на закрытость канала
			io.WriteString(os.Stdout, fmt.Sprintf("Worker %v channel %T closed\n", id, ch))
			return
		}
		io.WriteString(os.Stdout, fmt.Sprintf("Worker %v readed: '%v'\n", id, value))
	}
}

func main() {
	var N int
	fmt.Println("Enter N")
	fmt.Scan(&N)
	ch := make(chan int)
	for i := 0; i < N; i++ {
		go read(i, ch)
	}
	sigChan := make(chan os.Signal, 1)     //создание канала, отслеживающего Ctrl+C
	signal.Notify(sigChan, syscall.SIGINT) // уведомление о прерывании в канал

	func() {
		var i int
		for {
			i++
			select {
			case <-sigChan:
				return
			default:
				time.Sleep(time.Second)
				ch <- i
			}
		}
	}()
	close(ch) //закрытие канала, для завершения горутин
	close(sigChan)
	fmt.Print("stop")
	time.Sleep(time.Second * 2)
}
