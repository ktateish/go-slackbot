package brain

import (
	"fmt"
)

type OnMemoryBrain struct {
	db map[string][]byte
}

func NewOnMemoryBrain() *OnMemoryBrain {
	return &OnMemoryBrain{
		db: make(map[string][]byte),
	}
}

func (br *OnMemoryBrain) Load(key string) ([]byte, error) {
	v, ok := br.db[key]
	if !ok {
		return nil, fmt.Errorf("key='%s': %w", key, ErrNotFound)
	}
	return v, nil
}

func (br *OnMemoryBrain) Save(key string, val []byte) error {
	br.db[key] = val
	return nil
}
