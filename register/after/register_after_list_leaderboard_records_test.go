package after

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterAfterListLeaderboardRecords(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		out    *api.LeaderboardRecordList
		in     *api.ListLeaderboardRecordsRequest
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
			if err := RegisterAfterListLeaderboardRecords(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.out, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("RegisterAfterListLeaderboardRecords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
