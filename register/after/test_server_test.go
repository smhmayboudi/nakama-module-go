package after

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	u "github.com/smhmayboudi/nakama-modules-go/util"
)

func NewServer(t *testing.T) *httptest.Server {
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

	u.NewConfig(context.Background(), &TestLogger{})
	u.ModuleConfig.RedpandaURL = fmt.Sprintf("%s/topics/nakama", server.URL)

	return server
}
