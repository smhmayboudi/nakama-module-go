package util

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRedpandaSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/topics/nakama" {
			t.Errorf("Expected to request '/topics/nakama', got: %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/vnd.kafka.json.v2+json" {
			t.Errorf("Expected Content-Type: application/vnd.kafka.json.v2+json header, got: %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"offsets": [{"partition": 1, "offset": 0},{"partition": 2, "offset": 0},{"partition": 0, "offset": 0}]}`))
	}))
	defer server.Close()

	NewConfig(context.Background(), &TestLogger{})
	AppConfig.RedpandaURL = fmt.Sprintf("%s/topics/nakama", server.URL)

	type args struct {
		ctx     context.Context
		logger  runtime.Logger
		payload map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RedpandaSend",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				payload: map[string]interface{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RedpandaSend(tt.args.ctx, tt.args.logger, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("RedpandaSend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
