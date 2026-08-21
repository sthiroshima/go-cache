package internal

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Flush() {
	s.data = make(map[string]string)
}

func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		val = "nil"
	}

	return val, ok
}

func (s *Storage) Delete(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return 0
	}

	delete(s.data, key)

	return 1
}

func (s *Storage) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.data[key]; ok {
		return true
	}

	return false
}

func (s *Storage) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0)

	for k, _ := range s.data {
		keys = append(keys, k)
	}

	return keys
}

func (s *Storage) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}
