package after

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterAfterUnlinkFacebookInstantGame(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		in     *api.AccountFacebookInstantGame
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterAfterUnlinkFacebookInstantGame(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("RegisterAfterUnlinkFacebookInstantGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
