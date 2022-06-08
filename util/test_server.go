package util

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewServer(t *testing.T) (*httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/topics/nakama" {
			t.Errorf("Expected to request '/topics/nakama', got: %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/vnd.kafka.json.v2+json" {
			t.Errorf("Expected Content-Type: application/vnd.kafka.json.v2+json header, got: %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"offsets": [{"partition": 0, "offset": 0}]}`))
	}))

	NewConfig(context.Background(), &TestLogger{})
	ModuleConfig.RedpandaURL = fmt.Sprintf("%s/topics/nakama", server.URL)

	return server, nil
}
