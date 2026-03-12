package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestMyWaitGroup_Wait_BlocksUntilAllDone(t *testing.T) {
	wg := NewMyWaitGroup()
	var counter int64

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 5 {
		t.Errorf("Wait вернулся раньше времени: counter=%d, want 5", counter)
	}
}

func TestMyWaitGroup_Wait_ReturnsAfterDone(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(1)

	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Error("timeout: Wait не вернулся после Done")
	}
}

func TestMyWaitGroup_DoneBeforeWait_DoesNotBlock(t *testing.T) {
	// Done вызван раньше Wait — благодаря буферу 1 не блокируется
	wg := NewMyWaitGroup()
	wg.Add(1)
	wg.Done() // пишет в буфер, не блокируется

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Error("Wait заблокировался хотя счётчик уже 0")
	}
}

func TestMyWaitGroup_AddCalledMultipleTimes(t *testing.T) {
	wg := NewMyWaitGroup()
	var counter int64

	wg.Add(2)
	wg.Add(3)

	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 5 {
		t.Errorf("counter=%d, want 5", counter)
	}
}

func TestMyWaitGroup_MultipleGoroutines(t *testing.T) {
	wg := NewMyWaitGroup()
	const n = 100
	var counter int64

	wg.Add(int64(n))
	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != n {
		t.Errorf("counter=%d, want %d", counter, n)
	}
}

func TestCountTo_CallsDone(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(1)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	go CountTo(10, wg)

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Error("CountTo не вызвал Done")
	}
}

func TestCountTo_ZeroN(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(1)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	go CountTo(0, wg)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Error("CountTo(0) не завершился")
	}
}

func TestCountTo_MultipleGoroutines(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(5)

	go CountTo(11, wg)
	go CountTo(23, wg)
	go CountTo(4, wg)
	go CountTo(18, wg)
	go CountTo(32, wg)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Error("timeout: не все горутины завершились")
	}
}
