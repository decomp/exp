package bin_test

import (
	"encoding"
	"flag"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir/metadata"
)

var (
	_ flag.Value               = (*bin.Address)(nil)
	_ encoding.TextUnmarshaler = (*bin.Address)(nil)
	_ metadata.Unmarshaler     = (*bin.Address)(nil)
)
