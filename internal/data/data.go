package data

import (
	"runtime"
)

const AppName string = "balafetch"

const OS string = runtime.GOOS
const ARCH string = runtime.GOARCH

var Cards_categories = []string{
	"jokers",
	"tarot cards",
	"planet cards",
	"spectral cards",
	"vouchers",
}