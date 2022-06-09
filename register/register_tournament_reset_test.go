package register

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterTournamentReset(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		tournament *api.Tournament
		end        int64
		reset      int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RegisterTournamentReset",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         &TestNakamaModule{},
				tournament: &api.Tournament{},
				end:        0,
				reset:      0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterTournamentReset(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.tournament, tt.args.end, tt.args.reset); (err != nil) != tt.wantErr {
				t.Errorf("RegisterTournamentReset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
