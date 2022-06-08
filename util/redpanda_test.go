package util

import (
	"context"
	"fmt"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRedpandaSend(t *testing.T) {
	server, _ := NewServer(t)
	defer server.Close()

	NewConfig(context.Background(), &TestLogger{})
	ModuleConfig.RedpandaURL = fmt.Sprintf("%s/topics/nakama", server.URL)

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
