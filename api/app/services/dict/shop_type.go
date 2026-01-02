package dict

import "sync"

type ShopTypeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewShopTypeMap() *ShopTypeMap {
	return &ShopTypeMap{
		mu:   sync.RWMutex{},
		data: make(map[string]string),
	}
}

func (s *ShopTypeMap) Set(key string, value string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return s.data[key]
}
func (s *ShopTypeMap) Len() int {
	return len(s.data)
}
func (s *ShopTypeMap) GetData() *map[string]string {
	return &s.data
}
