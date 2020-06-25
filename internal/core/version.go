package core

import (
	"fmt"
	"runtime"
)

const (
	name    = "TelegramBot"
	version = "0.1.0"
)

var UserAgent = fmt.Sprintf("%s/%s (%s %s/%s)", name, version,
	runtime.Version(), runtime.GOOS, runtime.GOARCH)
