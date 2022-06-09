package before

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterBeforeSessionRefresh(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		in     *api.SessionRefreshRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *api.SessionRefreshRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterBeforeSessionRefresh(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterBeforeSessionRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterBeforeSessionRefresh() = %v, want %v", got, tt.want)
			}
		})
	}
}
