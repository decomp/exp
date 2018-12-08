package bin_test

import (
	"encoding"
	"flag"

	"github.com/decomp/exp/bin"
)

var (
	_ flag.Value               = (*bin.Address)(nil)
	_ encoding.TextUnmarshaler = (*bin.Address)(nil)
)
