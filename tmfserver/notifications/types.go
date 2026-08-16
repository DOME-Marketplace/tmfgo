package notifications

import (
	"log/slog"
	"slices"
	"time"
)

// Subscription represents a hub subscription created by a client for a given API family.
type Subscription struct {
	ID         string            `json:"id"`
	APIFamily  string            `json:"apiFamily"`
	Callback   string            `json:"callback"`
	EventTypes []string          `json:"eventTypes,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Query      string            `json:"query,omitempty"`
	CreatedAt  time.Time         `json:"createdAt"`
}

// Store abstracts persistence for subscriptions.
type Store interface {
	AddSubscription(sub *Subscription) error
	DeleteSubscription(apiFamily, id string) error
	GetSubscription(apiFamily, id string) (*Subscription, error)
	ListSubscriptionsByAPIFamily(apiFamily string) ([]*Subscription, error)
}

// DeliveryClient abstracts delivering events to subscriber callbacks.
type DeliveryClient interface {
	Deliver(sub *Subscription, payload any) error
}

// Manager coordinates subscriptions and event publishing.
type Manager struct {
	store   Store
	deliver DeliveryClient
}

func NewManager(store Store, deliver DeliveryClient) *Manager {
	return &Manager{store: store, deliver: deliver}
}

// CreateSubscription adds a new subscription for an API family.
func (m *Manager) CreateSubscription(apiFamily string, sub *Subscription) (*Subscription, error) {
	sub.APIFamily = apiFamily
	if err := m.store.AddSubscription(sub); err != nil {
		return nil, err
	}
	return sub, nil
}

// DeleteSubscription removes a subscription.
func (m *Manager) DeleteSubscription(apiFamily, id string) error {
	return m.store.DeleteSubscription(apiFamily, id)
}

// PublishEvent delivers the provided payload to all matching subscribers of the API family.
// Filtering by eventType is applied if the subscription specifies EventTypes.
func (m *Manager) PublishEvent(apiFamily, eventType string, payload any) {
	slog.Debug("generating event", "apiFamily", apiFamily, "eventType", eventType)
	subs, err := m.store.ListSubscriptionsByAPIFamily(apiFamily)
	if err != nil {
		return
	}
	for _, sub := range subs {
		if len(sub.EventTypes) > 0 {
			match := slices.Contains(sub.EventTypes, eventType)
			if !match {
				continue
			}
		}
		// Deliver asynchronously, ignoring individual errors (logged inside client)
		slog.Debug("deliver even asynchronously", "sub.id", sub.ID)
		go m.deliver.Deliver(sub, payload)
	}
}
