package bin

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

// Address represents a virtual address, which may be specified in hexadecimal
// notation. It implements the flag.Value, encoding.TextUnmarshaler and
// metadata.Unmarshaler interfaces.
type Address uint64

// String returns the hexadecimal string representation of v.
func (v Address) String() string {
	return fmt.Sprintf("0x%X", uint64(v))
}

// Set sets v to the numberic value represented by s.
func (v *Address) Set(s string) error {
	x, err := ParseUint64(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*v = Address(x)
	return nil
}

// UnmarshalText unmarshals the text into v.
func (v *Address) UnmarshalText(text []byte) error {
	return v.Set(string(text))
}

// MarshalText returns the textual representation of v.
func (v Address) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// InsertAddr inserts the given address within the sorted slice of addresses.
//
// pre-condition: addrs must be sorted in ascending order.
func InsertAddr(addrs []Address, addr Address) []Address {
	less := func(i int) bool {
		return addr <= addrs[i]
	}
	index := sort.Search(len(addrs), less)
	if index < len(addrs) && addrs[index] == addr {
		// addr is already present.
		return addrs
	}
	// addr is not present, insert at index.
	a := append(addrs[:index:index], addr)
	return append(a, addrs[index:]...)
}

// Addresses implements the sort.Sort interface, sorting addresses in ascending
// order.
type Addresses []Address

func (as Addresses) Len() int           { return len(as) }
func (as Addresses) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }
func (as Addresses) Less(i, j int) bool { return as[i] < as[j] }
