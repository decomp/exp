package bin_test

import (
	"encoding/json"
	"flag"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir/metadata"
)

var (
	_ flag.Value           = (*bin.Address)(nil)
	_ json.Unmarshaler     = (*bin.Address)(nil)
	_ metadata.Unmarshaler = (*bin.Address)(nil)
)
