package register

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterMatchmakerMatched(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx     context.Context
		logger  runtime.Logger
		db      *sql.DB
		nk      runtime.NakamaModule
		entries []runtime.MatchmakerEntry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestRegisterMatchmakerMatched",
			args: args{
				ctx:     context.Background(),
				logger:  &TestLogger{},
				db:      &sql.DB{},
				nk:      &TestNakamaModule{},
				entries: []runtime.MatchmakerEntry{},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterMatchmakerMatched(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterMatchmakerMatched() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RegisterMatchmakerMatched() = %v, want %v", got, tt.want)
			}
		})
	}
}
