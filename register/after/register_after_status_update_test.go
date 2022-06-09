package after

import (
	"context"
	"database/sql"
	"testing"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRegisterAfterStatusUpdate(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		out    *rtapi.Envelope
		in     *rtapi.Envelope
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
			if err := RegisterAfterStatusUpdate(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.out, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("RegisterAfterStatusUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
