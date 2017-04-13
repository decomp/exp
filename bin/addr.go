package bin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// An Address represents a virtual address; it implements the flag.Value,
// json.Unmarshaler and metadata.Unmarshaler interface.
type Address uint64

// String returns the hexadecimal string representation of the address.
func (v Address) String() string {
	return fmt.Sprintf("0x%X", uint64(v))
}

// Set sets v to the numberic value represented by s.
func (v *Address) Set(s string) error {
	base := 10
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
		base = 16
	}
	x, err := strconv.ParseUint(s, base, 64)
	if err != nil {
		return errors.WithStack(err)
	}
	*v = Address(x)
	return nil
}

// UnmarshalJSON unmarshals the data into v.
func (v *Address) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return errors.WithStack(err)
	}
	return v.Set(s)
}

// MarshalText returns the textual representation of v.
func (v Address) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalMetadata unmarshals the metadata node into v.
func (v *Address) UnmarshalMetadata(node metadata.Node) error {
	md, ok := node.(*metadata.Metadata)
	if !ok {
		return errors.Errorf("invalid metadata type; expected *metadata.Metadata, got %T", node)
	}
	if len(md.Nodes) != 1 {
		return errors.Errorf("invalid number of metadata nodes; expected 1, got %d", len(md.Nodes))
	}
	n := md.Nodes[0]
	s, ok := n.(*metadata.String)
	if !ok {
		return errors.Errorf("invalid metadata string type; expected *metadata.String, got %T", n)
	}
	return v.Set(s.Val)
}

// Addresses implements the sort.Sort interface, sorting addresses in ascending
// order.
type Addresses []Address

func (as Addresses) Len() int           { return len(as) }
func (as Addresses) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }
func (as Addresses) Less(i, j int) bool { return as[i] < as[j] }
