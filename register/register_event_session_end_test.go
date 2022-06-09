package register

import (
	"context"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterEventSessionEnd(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		evt    *api.Event
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "RegisterEventSessionEnd",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
				evt:    &api.Event{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterEventSessionEnd(tt.args.ctx, tt.args.logger, tt.args.evt)
		})
	}
}
