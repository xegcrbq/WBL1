package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
type counter struct {
	value int
	mx    sync.Mutex
	wg    sync.WaitGroup
}

func (c *counter) increment() {
	c.mx.Lock() //блокируем переменную, для других потоков, чтобы избежать потери данных
	c.value++   //инкреметируем переменную
	c.mx.Unlock()
	c.wg.Done()
}

func main() {

	c := counter{0, sync.Mutex{}, sync.WaitGroup{}}
	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		go c.increment()
	}
	c.wg.Wait()
	fmt.Println(c.value)
}
