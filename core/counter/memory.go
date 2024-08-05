package counter

import (
	"sync"
	"time"
)

const window = int64(60)

// MemoryCounter keeps track of request counts within a moving window of time.
type MemoryCounter struct {
	lock          sync.RWMutex  // lock to protect concurrent access
	LastInsertion int64         // timestamp of the last insertion
	StartupTime   int64         // timestamp of when the server started
	AccessCounts  [window]int64 // array to store request counts for each second in the window
}

// NewMemoryCounter initializes and returns a new instance of MemoryCounter.
func NewMemoryCounter() *MemoryCounter {
	return &MemoryCounter{
		LastInsertion: 0,
		StartupTime:   time.Now().Unix(),
		AccessCounts:  [window]int64{},
	}
}


// Access registers an access at the given timestamp (in seconds).
func (memCounter *MemoryCounter) Access(seconds int64) {
	memCounter.lock.Lock()
	defer memCounter.lock.Unlock()

	timeFromStart := seconds - memCounter.StartupTime
	if timeFromStart < window {
		memCounter.AccessCounts[seconds%window]++
		memCounter.LastInsertion = seconds
		return
	}

	preWindowEnd := seconds % window
	cellsToRemove := seconds - memCounter.LastInsertion

	if cellsToRemove >= window {
		memCounter.AccessCounts = [window]int64{}
	} else {
		for i := preWindowEnd; i > preWindowEnd-cellsToRemove; i-- {
			if i < 0 {
				memCounter.AccessCounts[i+window] = 0
			} else {
				memCounter.AccessCounts[i] = 0
			}
		}
	}

	memCounter.AccessCounts[seconds%window]++
	memCounter.LastInsertion = seconds
}

// Register returns the total number of requests counted in the moving window.
func (memCounter *MemoryCounter) Register() int64 {
	memCounter.lock.RLock()
	defer memCounter.lock.RUnlock()

	var total int64
	for _, count := range memCounter.AccessCounts {
		total += count
	}
	return total
}
