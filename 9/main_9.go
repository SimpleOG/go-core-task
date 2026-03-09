package main

import (
	"fmt"
	"math"
	"math/rand"
)

func NumConv(in chan uint8, out chan float64) {
	go func() {
		defer close(out)
		for n := range in {
			out <- math.Pow(float64(n), 3)
		}
	}()
}
func main() {
	in := make(chan uint8)
	out := make(chan float64)
	NumConv(in, out)
	go func() {
		//канал закроется когда все значения запишуться
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- uint8(rand.Intn(90))
		}
	}()
	for v := range out {
		fmt.Println(v)
	}
}
