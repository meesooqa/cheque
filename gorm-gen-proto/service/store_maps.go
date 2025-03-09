package service

import (
	"encoding/json"
	"os"
)

// Store stores "key => value" data
type Store struct {
	Data map[string]string
}

func NewStore(filename string) (*Store, error) {
	s := &Store{Data: make(map[string]string)}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &s.Data)
	if err != nil {
		return nil, err
	}
	return s, nil
}
