package main

import (
	"math"
	"testing"
	"time"
)

func collectFromOut(t *testing.T, out chan float64, timeout time.Duration) []float64 {
	t.Helper()
	result := make([]float64, 0)
	timer := time.After(timeout)
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return result
			}
			result = append(result, v)
		case <-timer:
			t.Fatal("timeout: out канал не закрылся, возможен дедлок")
			return nil
		}
	}
}

func runPipeline(t *testing.T, inputs []uint8) []float64 {
	t.Helper()
	in := make(chan uint8)
	out := make(chan float64)
	NumConv(in, out)

	go func() {
		defer close(in)
		for _, v := range inputs {
			in <- v
		}
	}()

	return collectFromOut(t, out, 3*time.Second)
}

// ===================== NumConv =====================

func TestNumConv_CubesValue(t *testing.T) {
	got := runPipeline(t, []uint8{3})
	want := math.Pow(3, 3) // 27

	if len(got) != 1 || got[0] != want {
		t.Errorf("NumConv(3) = %v, want %v", got, want)
	}
}

func TestNumConv_ZeroValue(t *testing.T) {
	got := runPipeline(t, []uint8{0})

	if len(got) != 1 || got[0] != 0 {
		t.Errorf("NumConv(0) = %v, want [0]", got)
	}
}

func TestNumConv_MaxUint8(t *testing.T) {
	// uint8 max = 255, 255^3 = 16581375
	got := runPipeline(t, []uint8{255})
	want := math.Pow(255, 3)

	if len(got) != 1 || got[0] != want {
		t.Errorf("NumConv(255) = %v, want %v", got, want)
	}
}

func TestNumConv_PreservesOrder(t *testing.T) {
	// Конвейер должен сохранять порядок: 1^3=1, 2^3=8, 3^3=27
	got := runPipeline(t, []uint8{1, 2, 3})
	want := []float64{1, 8, 27}

	if len(got) != len(want) {
		t.Fatalf("len(got)=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d]=%v, want %v", i, got[i], want[i])
		}
	}
}

func TestNumConv_AllValuesProcessed(t *testing.T) {
	inputs := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := runPipeline(t, inputs)

	if len(got) != len(inputs) {
		t.Fatalf("получили %d значений, want %d", len(got), len(inputs))
	}
	for i, input := range inputs {
		want := math.Pow(float64(input), 3)
		if got[i] != want {
			t.Errorf("NumConv(%d) = %v, want %v", input, got[i], want)
		}
	}
}

func TestNumConv_OutClosesWhenInCloses(t *testing.T) {
	// После закрытия in канал out должен закрыться
	in := make(chan uint8)
	out := make(chan float64)
	NumConv(in, out)

	close(in)

	// collectFromOut упадёт по таймауту если out не закроется
	collectFromOut(t, out, 2*time.Second)
}

func TestNumConv_EmptyInput_OutClosesImmediately(t *testing.T) {
	got := runPipeline(t, []uint8{})

	if len(got) != 0 {
		t.Errorf("пустой вход: got %v, want []", got)
	}
}
