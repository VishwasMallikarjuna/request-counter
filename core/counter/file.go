package counter

import (
	"encoding/json"
	"os"
	"sync"
)

// LoadFromFile loads the request counter data from a specified file.
func LoadFromFile(filename string) (*MemoryCounter, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var memCounter MemoryCounter
	if err := json.Unmarshal(data, &memCounter); err != nil {
		return nil, err
	}

	memCounter.lock = sync.RWMutex{}
	return &memCounter, nil
}

// SaveToFile saves the request counter data to a specified file.
func SaveToFile(filename string, memCounter *MemoryCounter) error {
	memCounter.lock.RLock()
	defer memCounter.lock.RUnlock()

	data, err := json.Marshal(memCounter)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
