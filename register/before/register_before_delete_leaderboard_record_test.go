package before

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterBeforeDeleteLeaderboardRecord(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		in     *api.DeleteLeaderboardRecordRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *api.DeleteLeaderboardRecordRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterBeforeDeleteLeaderboardRecord(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterBeforeDeleteLeaderboardRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterBeforeDeleteLeaderboardRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
