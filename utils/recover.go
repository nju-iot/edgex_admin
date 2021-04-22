package utils

import (
	"runtime/debug"

	"github.com/nju-iot/edgex_admin/logs"
)

// RecoverPanic ...
func RecoverPanic() {
	if x := recover(); x != nil {
		logs.Error("runtime panic: %v\n%v", x, string(debug.Stack()))
	}
}
