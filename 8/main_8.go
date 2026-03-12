package main

import (
	"fmt"
	"math"
	"sync/atomic"
)

type MyWaitGroup struct {
	count int64 // Для реализации аналога wg.Add\Done
	done  chan struct{}
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		done: make(chan struct{}, 1),
	}
}
func (wg *MyWaitGroup) Add(n int64) {
	//атомарно увеличиваем счетчик
	atomic.AddInt64(&wg.count, n)
}
func (wg *MyWaitGroup) Done() {
	//Нужно одновременно уменьшить счетчик и проверить
	// что счетчик не стал 0
	if atomic.AddInt64(&wg.count, -1) == 0 {
		//если счетчик стал 0, значит больше ничего ждать не нужно
		wg.done <- struct{}{}
	}
}


func (wg *MyWaitGroup) Wait() {
	//проверка, что счетчик не 0
	if atomic.LoadInt64(&wg.count) == 0 {
		return
	}

	<-wg.done
}

// функция для проверки работы кастомной wg
func CountTo(n int, wg *MyWaitGroup) {
	var res float64
	for i := 0; i < n; i++ {
		res += math.Pow(float64(i), 3)
	}
	fmt.Println(res)
	wg.Done()

}
func main() {
	wg := NewMyWaitGroup()
	wg.Add(5)
	go CountTo(11, wg)
	go CountTo(23, wg)
	go CountTo(4, wg)
	go CountTo(18, wg)
	go CountTo(32, wg)
	wg.Wait()

}
