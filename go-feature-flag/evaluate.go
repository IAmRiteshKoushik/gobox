package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:8080/evaluate/v1/boolean"

type EvaluationRequest struct {
	NamespaceKey string            `json:"namespaceKey"`
	FlagKey      string            `json:"flagKey"`
	EntityID     string            `json:"entityID"`
	Context      map[string]string `json:"context"`
}

type EvaluationResponse struct {
	Enabled bool `json:"enabled"`
}

func IsEnabled(ctx context.Context, flag string) (bool, error) {
	body, err := json.Marshal(EvaluationRequest{
		NamespaceKey: "default",
		FlagKey:      flag,
		EntityID:     "anonymous",
		Context:      map[string]string{},
	})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("flipt returned %d: %s", resp.StatusCode, data)
	}

	var result EvaluationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Enabled, nil
}
