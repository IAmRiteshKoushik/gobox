package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type FeatureStore struct {
	mu     sync.RWMutex
	flags  map[string]bool
	client *http.Client
	url    string
}

func NewFeatureStore(url string) *FeatureStore {
	return &FeatureStore{
		flags:  make(map[string]bool),
		client: http.DefaultClient,
		url:    url,
	}
}

func (s *FeatureStore) Reload(ctx context.Context) error {
	log.Println("Reloading flags from Flipt...")

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		s.url+"/api/v1/namespaces/default/flags",
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Flipt-Environment", "default")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("flipt returned %d", resp.StatusCode)
	}

	var result struct {
		Flags []struct {
			Key     string `json:"key"`
			Enabled bool   `json:"enabled"`
		} `json:"flags"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	next := make(map[string]bool)

	for _, flag := range result.Flags {
		next[flag.Key] = flag.Enabled
	}

	s.mu.Lock()
	s.flags = next
	log.Printf("%+v", s.flags)
	s.mu.Unlock()

	return nil
}

func (s *FeatureStore) StartPolling(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				if err := s.Reload(ctx); err != nil {
					log.Printf("failed to reload flags: %v", err)
				}
				log.Println("flags reloaded successfully")
			}
		}
	}()
}

func (s *FeatureStore) Enabled(flag string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.flags[flag]
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := NewFeatureStore("http://localhost:8080")

	if err := store.Reload(ctx); err != nil {
		log.Fatal(err)
	}

	// Refresh every 10 seconds.
	store.StartPolling(ctx, 5*time.Second)

	http.HandleFunc("/feature", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "new-login = %v\n", store.Enabled("new-login"))
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()

		fmt.Fprintf(w, "%#v\n", store.flags)
	})

	log.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
