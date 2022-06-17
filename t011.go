package main

import (
	"fmt"
	"sync"
)

//11.	Что выведет данная программа и почему?
//в начале 5 цифр в случайном порядке, затем дедлок
//в функции создаётся новый экземпляр WaitGroup не связанный с изначально объявленной
func main() {
	//правильный вариант

	//wg := &sync.WaitGroup{}
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(wg *sync.WaitGroup, j int) {
	//		fmt.Println(j)
	//		wg.Done()
	//	}(wg, i)
	//}
	//wg.Wait()
	//fmt.Println("exit")

	//вариант из примера

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, j int) {
			fmt.Println(j)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
