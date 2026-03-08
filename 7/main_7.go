package main

import (
	"fmt"
	"sync"
	"time"
)

func MergeChannels(channels ...chan int) chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))
		for _, c := range channels {
			go func(ch chan int) {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

// Заполнние канала данными
func FillChannel(values ...int) chan int {
	newChan := make(chan int)
	go func() {
		for _, v := range values {
			newChan <- v
			time.Sleep(1 * time.Second)
		}
		close(newChan)
	}()
	return newChan
}
func main() {
	//Напишите программу на Go, которая сливает N каналов в один.

	a := FillChannel(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := FillChannel(11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	c := FillChannel(21, 22, 23, 24, 25, 26, 27, 28, 29, 30)
	newChan := MergeChannels(a, b, c)
	for v := range newChan {
		fmt.Println(v)
	}
}
