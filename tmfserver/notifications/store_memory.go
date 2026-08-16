package notifications

import (
	"errors"
	"sync"
	"time"
)

var (
	errNotFound = errors.New("subscription not found")
)

// memoryStore is an in-memory implementation of Store.
type memoryStore struct {
	mu   sync.RWMutex
	data map[string]map[string]*Subscription // apiFamily -> id -> sub
}

func NewMemoryStore() Store {
	return &memoryStore{data: make(map[string]map[string]*Subscription)}
}

func (s *memoryStore) AddSubscription(sub *Subscription) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if sub.CreatedAt.IsZero() {
		sub.CreatedAt = time.Now()
	}
	if _, ok := s.data[sub.APIFamily]; !ok {
		s.data[sub.APIFamily] = make(map[string]*Subscription)
	}
	s.data[sub.APIFamily][sub.ID] = sub
	return nil
}

func (s *memoryStore) DeleteSubscription(apiFamily, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	family, ok := s.data[apiFamily]
	if !ok {
		return errNotFound
	}
	if _, ok := family[id]; !ok {
		return errNotFound
	}
	delete(family, id)
	return nil
}

func (s *memoryStore) GetSubscription(apiFamily, id string) (*Subscription, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	family, ok := s.data[apiFamily]
	if !ok {
		return nil, errNotFound
	}
	sub, ok := family[id]
	if !ok {
		return nil, errNotFound
	}
	return sub, nil
}

func (s *memoryStore) ListSubscriptionsByAPIFamily(apiFamily string) ([]*Subscription, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	family, ok := s.data[apiFamily]
	if !ok {
		return nil, nil
	}
	result := make([]*Subscription, 0, len(family))
	for _, sub := range family {
		result = append(result, sub)
	}
	return result, nil
}
