package util

import (
	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	InternalError = runtime.NewError("Internal Server Error", 500)
)
