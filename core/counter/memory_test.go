package counter

import (
	"sync"
	"testing"
	"time"
)

func TestNewMemoryCounter(t *testing.T) {
	counter := NewMemoryCounter()
	if counter == nil {
		t.Error("NewMemoryCounter() returned nil")
	}
	if counter.LastInsertion != 0 {
		t.Errorf("Expected LastInsertion to be 0, got %d", counter.LastInsertion)
	}
	if len(counter.AccessCounts) != int(window) {
		t.Errorf("Expected AccessCounts length to be %d, got %d", window, len(counter.AccessCounts))
	}
	for _, count := range counter.AccessCounts {
		if count != 0 {
			t.Error("Expected all AccessCounts to be 0")
		}
	}
}

func TestAccessWithinWindow(t *testing.T) {
	counter := NewMemoryCounter()
	startTime := counter.StartupTime
	counter.Access(startTime + 10)
	counter.Access(startTime + 10)
	if counter.AccessCounts[(startTime+10)%window] != 2 {
		t.Errorf("Expected AccessCounts at index %d to be 2, got %d", (startTime+10)%window, counter.AccessCounts[(startTime+10)%window])
	}
}

func TestAccessOutsideWindow(t *testing.T) {
	counter := NewMemoryCounter()
	startTime := counter.StartupTime
	counter.Access(startTime + 10)
	time.Sleep(1 * time.Second)
	counter.Access(startTime + window + 20)
	if counter.AccessCounts[(startTime+10)%window] != 0 {
		t.Errorf("Expected AccessCounts at index %d to be 0, got %d", (startTime+10)%window, counter.AccessCounts[(startTime+10)%window])
	}
	if counter.AccessCounts[(startTime+window+20)%window] != 1 {
		t.Errorf("Expected AccessCounts at index %d to be 1, got %d", (startTime+window+20)%window, counter.AccessCounts[(startTime+window+20)%window])
	}
}

func TestRegister(t *testing.T) {
	counter := NewMemoryCounter()
	startTime := counter.StartupTime
	counter.Access(startTime + 10)
	counter.Access(startTime + 20)
	counter.Access(startTime + 30)
	expectedTotal := int64(3)
	total := counter.Register()
	if total != expectedTotal {
		t.Errorf("Expected Register to return %d, got %d", expectedTotal, total)
	}
}

func TestRegisterAfterWindowExpiry(t *testing.T) {
	counter := NewMemoryCounter()
	startTime := counter.StartupTime
	counter.Access(startTime + 10)
	time.Sleep(1 * time.Second)
	counter.Access(startTime + window + 20)
	expectedTotal := int64(1)
	total := counter.Register()
	if total != expectedTotal {
		t.Errorf("Expected Register to return %d, got %d", expectedTotal, total)
	}
}

func TestConcurrency(t *testing.T) {
	counter := NewMemoryCounter()
	startTime := counter.StartupTime

	var wg sync.WaitGroup
	for i := int64(0); i < 100; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			counter.Access(startTime + i%window)
		}(i)
	}
	wg.Wait()

	total := counter.Register()
	expectedTotal := int64(100)
	if total != expectedTotal {
		t.Errorf("Expected Register to return %d, got %d", expectedTotal, total)
	}
}