// Package bin provides a uniform representation of binary executables.
package bin

// A File is a binary exectuable.
type File struct {
	// Machine architecture specifying the assembly instruction set.
	Arch Arch
	// Entry point of the executable.
	Entry Address
	// Segments of the exectuable.
	Segments []*Segment
	// Sections of the exectuable.
	Sections []*Section
}

// Arch represents the set of machine architectures.
type Arch uint

// Machine architectures.
const (
	// ArchX86_32 represents the 32-bit x86 machine architecture, as used by
	// Intel and AMD.
	ArchX86_32 Arch = 1 + iota
	// ArchX86_64 represents the 64-bit x86-64 machine architecture, as used by
	// Intel and AMD.
	ArchX86_64
)

// A Segment represents a continuous segment of memory.
type Segment struct {
	// Start address of the segment.
	Addr Address
	// Segment contents.
	Data []byte
	// Access permissions of the segment in memory.
	Perm Perm
}

// A Section represents a continuous section of memory.
type Section struct {
	// Section name.
	Name string
	// Start address of the section.
	Addr Address
	// Section contents.
	Data []byte
	// Access permissions of the section in memory.
	Perm Perm
}

// Perm specifies the access permissions of a segment or section in memory.
type Perm uint8

// Access permissions.
const (
	// PermR specifies that the memory is readable.
	PermR Perm = 1 << iota
	// PermW specifies that the memory is writeable.
	PermW
	// PermX specifies that the memory is executable.
	PermX
)

// String returns the string representation of the access permissions.
func (perm Perm) String() string {
	r := "-"
	if perm&PermR != 0 {
		r = "r"
	}
	w := "-"
	if perm&PermW != 0 {
		w = "w"
	}
	x := "-"
	if perm&PermX != 0 {
		x = "x"
	}
	return r + w + x
}
