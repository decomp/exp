package main

import (
	"fmt"

	"github.com/mewkiz/pkg/errutil"
	"rsc.io/x86/x86asm"
)

// convert converts the given binary excutable to equivalent C source code.
func convert(text []byte) error {
	panic("not yet implemented")
}

// convertFunc converts the function at the given offset in text to C source
// code.
func convertFunc(text []byte, offset int) error {
	for {
		inst, err := x86asm.Decode(text[offset:], 32)
		if err != nil {
			return errutil.Err(err)
		}
		fmt.Println(inst)
		offset += inst.Len
		if inst.Op == x86asm.RET {
			break
		}
	}
	return nil
}
