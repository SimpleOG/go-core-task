package main

import (
	"fmt"
	"math/rand"
)

func NumGenerator(done chan struct{}, limit int) chan int {
	out := make(chan int)
	go func() {
		//Канал закроется как только
		//вызыватель функции перестанет нуждаться в генерации чисел
		defer close(out)
		//бесконечная генерация до момента остановки
		for {
			//Слушаем 2 канала
			select {
			//генерация числа от 0 до limit в канал
			case out <- rand.Intn(limit):
			//случай завершения генерации
			case <-done:
				return

			}
		}
	}()
	//Сразу отдаю канал наружу
	return out
}
func main() {
	//Напишите генератор случайных чисел используя небуфферизированный канал.
	done := make(chan struct{})
	defer close(done)
	limit := 500
	randCh := NumGenerator(done, limit)
	for i := 0; i < 5; i++ {
		fmt.Printf("сгенерированно число %v \n", <-randCh)
	}
}
