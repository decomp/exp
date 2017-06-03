// Package elf provides access to Executable and Linkable Format (ELF) files.
package elf

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
)

// Register ELF format.
func init() {
	// Executable and Linkable Format (ELF)
	//
	//    7F 45 4C 46  |.ELF|
	const magic = "\x7FELF"
	bin.RegisterFormat("elf", magic, Parse)
}

// ParseFile parses the given ELF binary executable, reading from path.
func ParseFile(path string) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given ELF binary executable, reading from r.
//
// Users are responsible for closing r.
func Parse(r io.ReaderAt) (*bin.File, error) {
	// Open ELF file.
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse machine architecture.
	file := &bin.File{
		Exports: make(map[bin.Address]string),
	}
	switch f.Machine {
	case elf.EM_386:
		file.Arch = bin.ArchX86_32
	case elf.EM_X86_64:
		file.Arch = bin.ArchX86_64
	case elf.EM_PPC:
		file.Arch = bin.ArchPowerPC_32
	}

	// Parse entry address.
	file.Entry = bin.Address(f.Entry)

	// Parse sections.
	for _, s := range f.Sections {
		perm := parseSectFlags(s.Flags)
		data, err := s.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(data) == 0 {
			continue
		}
		sect := &bin.Section{
			Name: s.Name,
			Addr: bin.Address(s.Addr),
			Data: data,
			Perm: perm,
		}
		file.Sections = append(file.Sections, sect)
	}

	// Sort sections in ascending order.
	less := func(i, j int) bool {
		if file.Sections[i].Addr == file.Sections[j].Addr {
			if len(file.Sections[i].Data) > len(file.Sections[j].Data) {
				// prioritize longer sections with identical addresses.
				return true
			}
			return file.Sections[i].Name < file.Sections[j].Name
		}
		return file.Sections[i].Addr < file.Sections[j].Addr
	}
	sort.Slice(file.Sections, less)

	// Parse segments.
	var segments []*bin.Section
	for _, prog := range f.Progs {
		if prog.Type != elf.PT_LOAD {
			continue
		}
		r := prog.Open()
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		perm := parseProgFlags(prog.Flags)
		seg := &bin.Section{
			Addr: bin.Address(prog.Vaddr),
			Data: data,
			Perm: perm,
		}
		segments = append(segments, seg)
	}

	// Sort segments in ascending order.
	sort.Slice(segments, less)

	// Fix section permissions.
	if len(segments) > 0 {
		for _, sect := range file.Sections {
			for _, seg := range segments {
				end := seg.Addr + bin.Address(len(seg.Data))
				if seg.Addr <= sect.Addr && sect.Addr < end {
					if sect.Perm == 0 {
						sect.Perm = seg.Perm
					}
				}
			}
		}
	}

	// Append segments as sections.
	file.Sections = append(file.Sections, segments...)

	// Sort sections (and segments) in ascending order.
	sort.Slice(segments, less)

	// TODO: Parse imports.

	// Parse exports.
	symtab := f.Section(".symtab")
	strtab := f.Section(".strtab")
	if symtab != nil && strtab != nil {
		symtabData, err := symtab.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		strtabData, err := strtab.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		r := bytes.NewReader(symtabData)
		switch file.Arch.BitSize() {
		case 32:
			// Sym32 represents a 32-bit symbol descriptor.
			type Sym32 struct {
				// Index into the symbol string table.
				Name uint32
				// Value of the associated symbol. Depending on the context, this can
				// be an absolute value, an address, etc.
				Value uint32
				// Size in bytes; or 0 if the symbol has no size or an unknown size.
				Size uint32
				// Symbol type and binding information.
				Info uint8
				// Symbol visibility.
				Visibility SymVisibility
				// Section header table index relevant for the symbol.
				SectHdrIndex uint16
			}
			for {
				var sym Sym32
				if err := binary.Read(r, binary.LittleEndian, &sym); err != nil {
					if errors.Cause(err) == io.EOF {
						break
					}
					return nil, errors.WithStack(err)
				}
				//pretty.Println("sym:", sym)
				name := parseString(strtabData[sym.Name:])
				addr := bin.Address(sym.Value)
				typ := SymType(sym.Info & 0x0F)
				//bind := SymBind(sym.Info >> 4)
				// TODO: Remove debug output.
				//fmt.Println("name:", name)
				//fmt.Println("addr:", addr)
				//fmt.Println("size:", sym.Size)
				//fmt.Println("typ:", typ)
				//fmt.Println("bind:", bind)
				//fmt.Println("visibility:", sym.Visibility)
				//fmt.Println()
				if typ == SymTypeFunc {
					file.Exports[addr] = name
				}
			}
		case 64:
			// Sym64 represents a 64-bit symbol descriptor.
			type Sym64 struct {
				// Index into the symbol string table.
				Name uint32
				// Symbol type and binding information.
				Info uint8
				// Symbol visibility.
				Visibility SymVisibility
				// Section header table index relevant for the symbol.
				SectHdrIndex uint16
				// Value of the associated symbol. Depending on the context, this can
				// be an absolute value, an address, etc.
				Value uint64
				// Size in bytes; or 0 if the symbol has no size or an unknown size.
				Size uint64
			}
			for {
				var sym Sym64
				if err := binary.Read(r, binary.LittleEndian, &sym); err != nil {
					if errors.Cause(err) == io.EOF {
						break
					}
					return nil, errors.WithStack(err)
				}
				//pretty.Println("sym:", sym)
				name := parseString(strtabData[sym.Name:])
				addr := bin.Address(sym.Value)
				typ := SymType(sym.Info & 0x0F)
				//bind := SymBind(sym.Info >> 4)
				// TODO: Remove debug output.
				//fmt.Println("name:", name)
				//fmt.Println("addr:", addr)
				//fmt.Println("size:", sym.Size)
				//fmt.Println("typ:", typ)
				//fmt.Println("bind:", bind)
				//fmt.Println("visibility:", sym.Visibility)
				//fmt.Println()
				if typ == SymTypeFunc {
					file.Exports[addr] = name
				}
			}
		default:
			panic(fmt.Errorf("support for CPU bit size %d not yet implemented", file.Arch.BitSize()))
		}
	}

	return file, nil
}

// SymType specifies a symbol type.
type SymType uint8

// String returns the string representation of the symbol type.
func (typ SymType) String() string {
	m := map[SymType]string{
		SymTypeNone:    "none",
		SymTypeObject:  "object",
		SymTypeFunc:    "function",
		SymTypeSection: "section",
		SymTypeFile:    "file",
		SymTypeCommon:  "common",
		SymTypeOS0:     "OS 0",
		SymTypeOS1:     "OS 1",
		SymTypeOS2:     "OS 2",
		SymTypeProc0:   "processor 0",
		SymTypeProc1:   "processor 1",
		SymTypeProc2:   "processor 2",
	}
	if s, ok := m[typ]; ok {
		return s
	}
	panic(fmt.Errorf("support for symbol type %v not yet implemented", uint8(typ)))
}

// Symbol types.
const (
	// The symbol type is not specified.
	SymTypeNone SymType = 0
	// This symbol is associated with a data object, such as a variable, an
	// array, and so forth.
	SymTypeObject SymType = 1
	// This symbol is associated with a function or other executable code.
	SymTypeFunc SymType = 2
	// This symbol is associated with a section.
	SymTypeSection SymType = 3
	// Name of the source file associated with the object file
	SymTypeFile SymType = 4
	// This symbol labels an uninitialized common block.
	SymTypeCommon SymType = 5
	// Reserved for operating system-specific semantics.
	SymTypeOS0 SymType = 10
	// Reserved for operating system-specific semantics.
	SymTypeOS1 SymType = 11
	// Reserved for operating system-specific semantics.
	SymTypeOS2 SymType = 12
	// Reserved for processor-specific semantics.
	SymTypeProc0 SymType = 13
	// Reserved for processor-specific semantics.
	SymTypeProc1 SymType = 14
	// Reserved for processor-specific semantics.
	SymTypeProc2 SymType = 15
)

// SymBind specifies a symbol binding.
type SymBind uint8

// String returns the string representation of the symbol binding.
func (bind SymBind) String() string {
	m := map[SymBind]string{
		SymBindLocal:  "local",
		SymBindGlobal: "global",
		SymBindWeak:   "weak",
		SymBindOS0:    "OS 0",
		SymBindOS1:    "OS 1",
		SymBindOS2:    "OS 2",
		SymBindProc0:  "processor 0",
		SymBindProc1:  "processor 1",
		SymBindProc2:  "processor 2",
	}
	if s, ok := m[bind]; ok {
		return s
	}
	panic(fmt.Errorf("support for symbol binding %v not yet implemented", uint8(bind)))
}

// Symbol bindings.
const (
	// Local symbol.
	SymBindLocal SymBind = 0
	// Global symbol.
	SymBindGlobal SymBind = 1
	// Weak symbol.
	SymBindWeak SymBind = 2
	// Reserved for operating system-specific semantics.
	SymBindOS0 SymBind = 10
	// Reserved for operating system-specific semantics.
	SymBindOS1 SymBind = 11
	// Reserved for operating system-specific semantics.
	SymBindOS2 SymBind = 12
	// Reserved for processor-specific semantics.
	SymBindProc0 SymBind = 13
	// Reserved for processor-specific semantics.
	SymBindProc1 SymBind = 14
	// Reserved for processor-specific semantics.
	SymBindProc2 SymBind = 15
)

// SymVisibility specifies a symbol visibility.
type SymVisibility uint8

// String returns the string representation of the symbol binding.
func (v SymVisibility) String() string {
	m := map[SymVisibility]string{
		SymVisibilityDefault:   "default",
		SymVisibilityInternal:  "internal",
		SymVisibilityHidden:    "hidden",
		SymVisibilityProtected: "protected",
	}
	if s, ok := m[v]; ok {
		return s
	}
	panic(fmt.Errorf("support for symbol visibility %v not yet implemented", uint8(v)))
}

// Symbol visibility.
const (
	// Default symbol visiblity as specified by the symbol binding.
	SymVisibilityDefault SymVisibility = 0
	// Internal symbol visibility.
	SymVisibilityInternal SymVisibility = 1
	// Hidden symbol visibility.
	SymVisibilityHidden SymVisibility = 2
	// Protected symbol visibility.
	SymVisibilityProtected SymVisibility = 3
)

// parseSectFlags returns the memory access permissions represented by the given
// section header flags.
func parseSectFlags(flags elf.SectionFlag) bin.Perm {
	var perm bin.Perm
	if flags&elf.SHF_WRITE != 0 {
		perm |= bin.PermW
	}
	if flags&elf.SHF_EXECINSTR != 0 {
		perm |= bin.PermX
	}
	return perm
}

// parseProgFlags returns the memory access permissions represented by the given
// program header flags.
func parseProgFlags(flags elf.ProgFlag) bin.Perm {
	var perm bin.Perm
	if flags&elf.PF_R != 0 {
		perm |= bin.PermR
	}
	if flags&elf.PF_W != 0 {
		perm |= bin.PermW
	}
	if flags&elf.PF_X != 0 {
		perm |= bin.PermX
	}
	return perm
}

// ### [ Helper functions ] ####################################################

// parseString parses the NULL-terminated string in the given data.
func parseString(data []byte) string {
	pos := bytes.IndexByte(data, '\x00')
	if pos == -1 {
		panic(fmt.Errorf("unable to locate NULL-terminated string in % 02X", data))
	}
	return string(data[:pos])
}
