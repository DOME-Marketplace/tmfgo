package notifications

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type httpDelivery struct {
	client *http.Client
}

func NewHTTPDelivery(timeout time.Duration) DeliveryClient {
	return &httpDelivery{client: &http.Client{Timeout: timeout}}
}

func (d *httpDelivery) Deliver(sub *Subscription, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var lastErr error
	for attempt := range 3 {
		req, err := http.NewRequest("POST", sub.Callback, bytes.NewReader(body))
		if err != nil {
			lastErr = err
			break
		}
		req.Header.Set("Content-Type", "application/json")
		// Pass through subscriber-provided auth header if present
		if token, ok := sub.Headers["x-auth-token"]; ok && token != "" {
			req.Header.Set("x-auth-token", token)
		}

		resp, err := d.client.Do(req)
		if err == nil && resp != nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			_ = resp.Body.Close()
			return nil
		}
		if resp != nil {
			_ = resp.Body.Close()
		}
		lastErr = err
		slog.Warn("notification delivery failed, will retry", slog.String("callback", sub.Callback), slog.Int("attempt", attempt+1))
		time.Sleep(time.Duration(1<<attempt) * 200 * time.Millisecond)
	}
	return lastErr
}
