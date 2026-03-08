package main

import (
	"sort"
	"testing"
	"time"
)

func collectWithTimeout(t *testing.T, ch chan int, timeout time.Duration) []int {
	t.Helper()
	result := make([]int, 0)
	timer := time.After(timeout)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return result
			}
			result = append(result, v)
		case <-timer:
			t.Fatalf("timeout: выходной канал не закрылся за %v, возможен дедлок", timeout)
			return nil
		}
	}
}

func TestMergeChannels_AllValuesReceived(t *testing.T) {
	a := FillChannel(1, 2, 3)
	b := FillChannel(4, 5, 6)
	c := FillChannel(7, 8, 9)

	out := MergeChannels(a, b, c)
	got := collectWithTimeout(t, out, 5*time.Second)

	if len(got) != 9 {
		t.Errorf("ожидали 9 элементов, получили %d: %v", len(got), got)
	}
}

func TestMergeChannels_NoValueLost(t *testing.T) {
	a := FillChannel(1, 2, 3)
	b := FillChannel(4, 5, 6)
	c := FillChannel(7, 8, 9)

	out := MergeChannels(a, b, c)
	got := collectWithTimeout(t, out, 5*time.Second)

	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if len(got) != len(want) {
		t.Fatalf("len(got)=%d, len(want)=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d]=%d, want %d; full got: %v", i, got[i], want[i], got)
		}
	}
}

func TestMergeChannels_OutputChannelCloses(t *testing.T) {
	a := FillChannel(1, 2)
	b := FillChannel(3, 4)

	out := MergeChannels(a, b)

	collectWithTimeout(t, out, 5*time.Second)
}

func TestMergeChannels_SingleChannel(t *testing.T) {
	a := FillChannel(10, 20, 30)

	out := MergeChannels(a)
	got := collectWithTimeout(t, out, 5*time.Second)

	sort.Ints(got)
	want := []int{10, 20, 30}

	if len(got) != len(want) {
		t.Fatalf("len(got)=%d, len(want)=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestMergeChannels_EmptyChannels(t *testing.T) {
	a := FillChannel()
	b := FillChannel()

	out := MergeChannels(a, b)
	got := collectWithTimeout(t, out, 3*time.Second)

	if len(got) != 0 {
		t.Errorf("ожидали пустой слайс, получили %v", got)
	}
}

func TestMergeChannels_ZeroChannels(t *testing.T) {
	out := MergeChannels()
	got := collectWithTimeout(t, out, 3*time.Second)

	if len(got) != 0 {
		t.Errorf("ожидали пустой слайс, получили %v", got)
	}
}

func TestMergeChannels_ManyChannels(t *testing.T) {
	const n = 10
	channels := make([]chan int, n)
	want := make([]int, 0, n*3)

	for i := 0; i < n; i++ {
		v1, v2, v3 := i*3+1, i*3+2, i*3+3
		channels[i] = FillChannel(v1, v2, v3)
		want = append(want, v1, v2, v3)
	}

	out := MergeChannels(channels...)
	got := collectWithTimeout(t, out, 10*time.Second)

	if len(got) != len(want) {
		t.Fatalf("len(got)=%d, len(want)=%d", len(got), len(want))
	}

	sort.Ints(got)
	sort.Ints(want)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestFillChannel_ReturnsAllValues(t *testing.T) {
	ch := FillChannel(5, 10, 15)
	got := collectWithTimeout(t, ch, 5*time.Second)

	want := []int{5, 10, 15}
	if len(got) != len(want) {
		t.Fatalf("len(got)=%d, len(want)=%d; got: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d]=%d, want %d", i, got[i], want[i])
		}
	}
}

func TestFillChannel_ClosesAfterAllValues(t *testing.T) {
	ch := FillChannel(1, 2, 3)
	collectWithTimeout(t, ch, 5*time.Second)
}

func TestFillChannel_Empty_ClosesImmediately(t *testing.T) {
	ch := FillChannel()
	got := collectWithTimeout(t, ch, 2*time.Second)

	if len(got) != 0 {
		t.Errorf("ожидали пустой слайс, получили %v", got)
	}
}
