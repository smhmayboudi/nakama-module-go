package before

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterBeforeMatchJoin(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		in     *rtapi.Envelope
	}
	tests := []struct {
		name    string
		args    args
		want    *rtapi.Envelope
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterBeforeMatchJoin(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterBeforeMatchJoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterBeforeMatchJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
