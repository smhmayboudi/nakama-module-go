package util

import (
	"context"
	"fmt"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRedpandaSend(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	NewConfig(context.Background(), &TestLogger{})

	type args struct {
		ctx     context.Context
		logger  runtime.Logger
		payload map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		init    func()
	}{
		{
			name: "RedpandaSend",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				payload: map[string]interface{}{},
			},
			wantErr: false,
			init: func() {
				ModuleConfig.RedpandaURL = fmt.Sprintf("%s/topics/nakama", server.URL)
			},
		},
		{
			name: "RedpandaSendWithError",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				payload: map[string]interface{}{"error": func() {}},
			},
			wantErr: true,
			init: func() {
			},
		},
		{
			name: "RedpandaSendWithError2",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				payload: map[string]interface{}{},
			},
			wantErr: true,
			init: func() {
				ModuleConfig.RedpandaURL = fmt.Sprintf(":%s/topics/nakama", server.URL)
			},
		},
		{
			name: "RedpandaSendWithError3",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				payload: map[string]interface{}{},
			},
			wantErr: true,
			init: func() {
				ModuleConfig.RedpandaURL = "*"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init()
			if err := RedpandaSend(tt.args.ctx, tt.args.logger, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("RedpandaSend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
