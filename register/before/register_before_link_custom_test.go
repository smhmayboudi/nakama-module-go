package before

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterBeforeLinkCustom(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		in     *api.AccountCustom
	}
	tests := []struct {
		name    string
		args    args
		want    *api.AccountCustom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterBeforeLinkCustom(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterBeforeLinkCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterBeforeLinkCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}
