package register

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterLeaderboardReset(t *testing.T) {
	server, _ := NewServer(t)
	defer server.Close()

	type args struct {
		ctx         context.Context
		logger      runtime.Logger
		db          *sql.DB
		nk          runtime.NakamaModule
		leaderboard *api.Leaderboard
		reset       int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RegisterLeaderboardReset",
			args: args{
				ctx:         context.Background(),
				logger:      &TestLogger{},
				db:          &sql.DB{},
				nk:          nil,
				leaderboard: &api.Leaderboard{},
				reset:       1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterLeaderboardReset(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.leaderboard, tt.args.reset); (err != nil) != tt.wantErr {
				t.Errorf("RegisterLeaderboardReset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
