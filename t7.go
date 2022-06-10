package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func write(m map[int]int, i int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	mutex.Lock() //блокирование мютекса на время записи
	m[i] = i
	mutex.Unlock() // разблокирование мютекса
	wg.Done()
}

func main() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}

	for i := 0; i < 5; i++ { //запуск горутин

		wg.Add(1)
		go write(m, i, &wg, &mutex)
	}

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println(m)
}
