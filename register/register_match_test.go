// https://github.com/heroiclabs/nakama/blob/master/sample_go_module/sample.go
// https://heroiclabs.com/docs/nakama/concepts/multiplayer/authoritative/
package register

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestMatch_MatchInit(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
		params map[string]interface{}
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
		want1 int
		want2 string
	}{
		{
			name: "Match_MatchInit",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
				db:     &sql.DB{},
				nk:     nil,
				params: map[string]interface{}{
					"debug": false,
				},
			},
			want: &MatchState{
				debug: false,
			},
			want1: 1,
			want2: "skill=100-150",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			got, got1, got2 := match.MatchInit(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchInit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Match.MatchInit() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Match.MatchInit() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMatch_MatchJoinAttempt(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		dispatcher runtime.MatchDispatcher
		tick       int64
		state      interface{}
		presence   runtime.Presence
		metadata   map[string]string
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
		want1 bool
		want2 string
	}{
		{
			name: "Match_MatchJoinAttempt",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
			},
			want: &MatchState{
				debug: false,
			},
			want1: true,
			want2: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			got, got1, got2 := match.MatchJoinAttempt(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.presence, tt.args.metadata)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchJoinAttempt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Match.MatchJoinAttempt() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Match.MatchJoinAttempt() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMatch_MatchJoin(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		dispatcher runtime.MatchDispatcher
		tick       int64
		state      interface{}
		presences  []runtime.Presence
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
	}{
		{
			name: "Match_MatchJoin",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
			},
			want: &MatchState{
				debug: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			if got := match.MatchJoin(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.presences); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_MatchLeave(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		dispatcher runtime.MatchDispatcher
		tick       int64
		state      interface{}
		presences  []runtime.Presence
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
	}{
		{
			name: "Match_MatchLeave",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
			},
			want: &MatchState{
				debug: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			if got := match.MatchLeave(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.presences); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchLeave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_MatchLoop(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		dispatcher runtime.MatchDispatcher
		tick       int64
		state      interface{}
		messages   []runtime.MatchData
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
	}{
		{
			name: "Match_MatchLoop",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
			},
			want: &MatchState{
				debug: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			if got := match.MatchLoop(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.messages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_MatchTerminate(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx          context.Context
		logger       runtime.Logger
		db           *sql.DB
		nk           runtime.NakamaModule
		dispatcher   runtime.MatchDispatcher
		tick         int64
		state        interface{}
		graceSeconds int
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
	}{
		{
			name: "Match_MatchTerminate",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
			},
			want: &MatchState{
				debug: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			if got := match.MatchTerminate(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.graceSeconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchTerminate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_MatchSignal(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx        context.Context
		logger     runtime.Logger
		db         *sql.DB
		nk         runtime.NakamaModule
		dispatcher runtime.MatchDispatcher
		tick       int64
		state      interface{}
		data       string
	}
	tests := []struct {
		name  string
		match *Match
		args  args
		want  interface{}
		want1 string
	}{
		{
			name: "Match_MatchSignal",
			args: args{
				ctx:        context.Background(),
				logger:     &TestLogger{},
				db:         &sql.DB{},
				nk:         nil,
				dispatcher: nil,
				tick:       0,
				state: &MatchState{
					debug: false,
				},
				data: "",
			},
			want: &MatchState{
				debug: false,
			},
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := &Match{}
			got, got1 := match.MatchSignal(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk, tt.args.dispatcher, tt.args.tick, tt.args.state, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MatchSignal() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Match.MatchSignal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegisterMatch(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	type args struct {
		ctx    context.Context
		logger runtime.Logger
		db     *sql.DB
		nk     runtime.NakamaModule
	}
	tests := []struct {
		name    string
		args    args
		want    runtime.Match
		wantErr bool
	}{
		{
			name: "RegisterMatch",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
				db:     &sql.DB{},
				nk:     nil,
			},
			want:    &Match{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterMatch(tt.args.ctx, tt.args.logger, tt.args.db, tt.args.nk)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
