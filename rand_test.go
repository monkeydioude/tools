package tools

import (
	"sync"
	"testing"
	"time"
)

func TestIGetDifferentIntegerOnSuccessiveCalls(t *testing.T) {
	assertDiff(t, RandUnixNano(50), RandUnixNano(50))
}

func TestIGetDifferentDurationOnSuccessiveCalls(t *testing.T) {
	assertDiff(t, RandDurationUnixNano(50), RandDurationUnixNano(50))
}

func TestIGetDifferentIntegerOnGoRoutinesCall(t *testing.T) {
	var wg sync.WaitGroup
	var a, b int64

	wg.Add(2)

	go func() {
		a = RandUnixNano(32)
		wg.Done()
	}()

	go func() {
		b = RandUnixNano(32)
		wg.Done()
	}()

	wg.Wait()
	assertDiff(t, a, b)
}

func TestIGetDifferentDurationOnGoRoutinesCall(t *testing.T) {
	var wg sync.WaitGroup
	var a, b time.Duration

	wg.Add(2)

	go func() {
		a = RandDurationUnixNano(32 * time.Second)
		wg.Done()
	}()

	go func() {
		b = RandDurationUnixNano(32 * time.Second)
		wg.Done()
	}()

	wg.Wait()
	assertDiff(t, a, b)
}
