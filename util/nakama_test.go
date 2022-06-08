package util

import (
	"context"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestNewContext(t *testing.T) {
	type args struct {
		ctx    context.Context
		logger runtime.Logger
	}
	tests := []struct {
		name string
		args args
		want *NakamaContext
	}{
		{
			name: "NewContext",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
			},
			want: &NakamaContext{
				ClientIp:       "",
				ClientPort:     "",
				Env:            map[string]string{},
				ExecutionMode:  "",
				Headers:        map[string]string{},
				Lang:           "",
				MatchId:        "",
				MatchLabel:     "",
				MatchNode:      "",
				MatchTickRate:  0,
				Node:           "",
				QueryParams:    map[string]string{},
				SessionId:      "",
				UserId:         "",
				UserSessionExp: 0,
				Username:       "",
				Vars:           map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContext(tt.args.ctx, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
