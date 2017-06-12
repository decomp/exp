package x86

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
)

// Contexts tracks the CPU context at various addresses of the executable.
type Contexts map[bin.Address]Context

// Context tracks the CPU context at a specific address of the executable.
type Context struct {
	// Register constraints.
	Regs map[Register]ValueContext `json:"regs"`
	// Instruction argument constraints.
	Args map[int]ValueContext `json:"args"`
}

// ValueContext defines constraints on a value used at a specific address.
//
// The following keys are defined.
//
//    Key          Type          Description
//
//    addr         bin.Address   virtual address.
//
//    min          int64         minimum value.
//    max          int64         maximum value.
//
//    Mem.offset   int64         memory reference offset.
//
//    symbol       string        symbol name.
type ValueContext map[string]Value

// Value represents a value at a specific address.
type Value struct {
	// underlying string representation of the value.
	s string
}

// String returns the string representation of v.
func (v Value) String() string {
	return v.s
}

// Set sets v to the value represented by s.
func (v *Value) Set(s string) error {
	v.s = s
	return nil
}

// UnmarshalText unmarshals the text into v.
func (v *Value) UnmarshalText(text []byte) error {
	if err := v.Set(string(text)); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// MarshalText returns the textual representation of v.
func (v Value) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// Addr returns the virtual address represented by v.
func (v Value) Addr() bin.Address {
	var addr bin.Address
	if err := addr.Set(v.s); err != nil {
		panic(fmt.Errorf("unable to parse value %q as virtual address; %v", v.s, err))
	}
	return addr
}

// Int64 returns the 64-bit signed integer represented by v.
func (v Value) Int64() int64 {
	x, err := strconv.ParseInt(v.s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse value %q as int64; %v", v.s, err))
	}
	return x
}

// Uint64 returns the 64-bit unsigned integer represented by v.
func (v Value) Uint64() uint64 {
	s := v.s
	base := 10
	if strings.HasPrefix(s, "0x") {
		s = s[len("0x"):]
	}
	x, err := strconv.ParseUint(s, base, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse value %q as uint64; %v", v.s, err))
	}
	return x
}
