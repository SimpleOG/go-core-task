package main

import (
	"testing"
	"time"
)

func TestNumGenerator_ChannelIsUnbuffered(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ch := NumGenerator(done, 100)

	if cap(ch) != 0 {
		t.Errorf("ожидали небуферизированный канал (cap=0), got cap=%d", cap(ch))
	}
}

func TestNumGenerator_ProducesValues(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ch := NumGenerator(done, 100)

	select {
	case _, ok := <-ch:
		if !ok {
			t.Error("канал закрылся раньше времени")
		}
	case <-time.After(time.Second):
		t.Error("timeout: генератор не выдал значение за 1 секунду")
	}
}

func TestNumGenerator_ProducesMultipleValues(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ch := NumGenerator(done, 100)

	const count = 10
	for i := 0; i < count; i++ {
		select {
		case <-ch:
		case <-time.After(time.Second):
			t.Fatalf("timeout на %d-м значении", i+1)
		}
	}
}

func TestNumGenerator_StopsWhenDoneClosed(t *testing.T) {
	done := make(chan struct{})
	ch := NumGenerator(done, 100)

	// убеждаемся что генератор работает
	<-ch
	<-ch

	close(done)

	// после close(done) канал out должен закрыться
	timer := time.After(time.Second)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				// канал закрылся — ожидаемое поведение
				return
			}
		case <-timer:
			t.Error("timeout: канал не закрылся после close(done)")
			return
		}
	}
}

func TestNumGenerator_ValuesInRange(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	const limit = 100
	ch := NumGenerator(done, limit)

	for i := 0; i < 50; i++ {
		select {
		case v := <-ch:
			if v < 0 || v >= limit {
				t.Errorf("значение %d вне диапазона [0, %d)", v, limit)
			}
		case <-time.After(time.Second):
			t.Fatalf("timeout на %d-м значении", i+1)
		}
	}
}

func TestNumGenerator_RespectsLimit(t *testing.T) {
	// Проверяем что limit действительно влияет на диапазон значений
	done := make(chan struct{})
	defer close(done)

	const limit = 5
	ch := NumGenerator(done, limit)

	for i := 0; i < 50; i++ {
		select {
		case v := <-ch:
			if v >= limit {
				t.Errorf("значение %d >= limit %d", v, limit)
			}
		case <-time.After(time.Second):
			t.Fatalf("timeout на %d-м значении", i+1)
		}
	}
}

func TestNumGenerator_ValuesAreRandom(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ch := NumGenerator(done, 500)

	const count = 20
	first := <-ch
	allSame := true
	for i := 1; i < count; i++ {
		select {
		case v := <-ch:
			if v != first {
				allSame = false
			}
		case <-time.After(time.Second):
			t.Fatalf("timeout на %d-м значении", i+1)
		}
	}

	if allSame {
		t.Errorf("все %d значений одинаковые (%d) — генератор не случайный?", count, first)
	}
}
