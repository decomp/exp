package bin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Uint64 represents a 64-bit unsigned integer, which may be specified in
// hexadecimal notation. It implements the flag.Value and
// encoding.TextUnmarshaler interfaces.
type Uint64 uint64

// String returns the hexadecimal string representation of v.
func (v Uint64) String() string {
	return fmt.Sprintf("0x%X", uint64(v))
}

// Set sets v to the numberic value represented by s.
func (v *Uint64) Set(s string) error {
	x, err := ParseUint64(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*v = x
	return nil
}

// UnmarshalText unmarshals the text into v.
func (v *Uint64) UnmarshalText(text []byte) error {
	return v.Set(string(text))
}

// MarshalText returns the textual representation of v.
func (v Uint64) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// ParseUint64 interprets the given string in base 10 or base 16 (if prefixed
// with `0x` or `0X`) and returns the corresponding value.
func ParseUint64(s string) (Uint64, error) {
	base := 10
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
		base = 16
	}
	x, err := strconv.ParseUint(s, base, 64)
	if err != nil {
		// Parse signed integer as fallback.
		y, err := strconv.ParseInt(s, base, 64)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		return Uint64(y), nil
	}
	return Uint64(x), nil
}
