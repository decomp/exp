// Package pe provides access to PE (Portable Executable) files.
package pe

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

// Register PE format.
func init() {
	// Portable Executable (PE) format.
	//
	//    4D 5A  |MZ|
	const magic = "MZ"
	bin.RegisterFormat("pe", magic, Parse)
}

// ParseFile parses the given PE binary executable, reading from path.
func ParseFile(path string) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given PE binary executable, reading from r.
//
// Users are responsible for closing r.
func Parse(r io.ReaderAt) (*bin.File, error) {
	// Open PE file.
	f, err := pe.NewFile(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse machine architecture.
	file := &bin.File{}
	switch f.FileHeader.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		file.Arch = bin.ArchX86_32
	case pe.IMAGE_FILE_MACHINE_AMD64:
		file.Arch = bin.ArchX86_64
	case pe.IMAGE_FILE_MACHINE_POWERPC:
		file.Arch = bin.ArchPowerPC_32
	default:
		panic(fmt.Errorf("support for machine architecture %v not yet implemented", f.FileHeader.Machine))
	}

	// Parse entry address.
	var (
		// Image base address.
		imageBase uint64
		// Import table RVA and size.
		itRVA  uint64
		itSize uint64
		// Import address table (IAT) RVA and size.
		iatRVA  uint64
		iatSize uint64
	)
	// Data directory indices.
	const (
		ImportTableIndex        = 1
		ImportAddressTableIndex = 12
	)
	switch opt := f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		file.Entry = bin.Address(opt.ImageBase + opt.AddressOfEntryPoint)
		imageBase = uint64(opt.ImageBase)
		itRVA = uint64(opt.DataDirectory[ImportTableIndex].VirtualAddress)
		itSize = uint64(opt.DataDirectory[ImportTableIndex].Size)
		iatRVA = uint64(opt.DataDirectory[ImportAddressTableIndex].VirtualAddress)
		iatSize = uint64(opt.DataDirectory[ImportAddressTableIndex].Size)
	case *pe.OptionalHeader64:
		file.Entry = bin.Address(opt.ImageBase) + bin.Address(opt.AddressOfEntryPoint)
		imageBase = uint64(opt.ImageBase)
		itRVA = uint64(opt.DataDirectory[ImportTableIndex].VirtualAddress)
		itSize = uint64(opt.DataDirectory[ImportTableIndex].Size)
		iatRVA = uint64(opt.DataDirectory[ImportAddressTableIndex].VirtualAddress)
		iatSize = uint64(opt.DataDirectory[ImportAddressTableIndex].Size)
	default:
		panic(fmt.Errorf("support for optional header type %T not yet implemented", opt))
	}

	// Parse sections.
	for _, s := range f.Sections {
		addr := bin.Address(imageBase) + bin.Address(s.VirtualAddress)
		data, err := s.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		perm := parsePerm(s.Characteristics)
		sect := &bin.Section{
			Name: s.Name,
			Addr: addr,
			Data: data,
			Perm: perm,
		}
		file.Sections = append(file.Sections, sect)
	}
	less := func(i, j int) bool {
		if file.Sections[i].Addr == file.Sections[j].Addr {
			return file.Sections[i].Name < file.Sections[j].Name
		}
		return file.Sections[i].Addr < file.Sections[j].Addr
	}
	sort.Slice(file.Sections, less)

	// Parse import address table (IAT).
	fmt.Println("iat")
	if iatSize != 0 {
		iatAddr := bin.Address(imageBase + iatRVA)
		fmt.Println("iat addr:", iatAddr)
		data := file.Data(iatAddr)
		data = data[:iatSize]
		fmt.Println(hex.Dump(data))
	}

	// Parse import table.
	fmt.Println("it")
	var imps []importDesc
	if itSize != 0 {
		itAddr := bin.Address(imageBase + itRVA)
		fmt.Println("it addr:", itAddr)
		data := file.Data(itAddr)
		data = data[:itSize]
		fmt.Println(hex.Dump(data))
		r := bytes.NewReader(data)
		zero := importDesc{}
		for {
			var imp importDesc
			if err := binary.Read(r, binary.LittleEndian, &imp); err != nil {
				return nil, errors.WithStack(err)
			}
			pretty.Println("imp:", imp)
			if imp == zero {
				break
			}
			imps = append(imps, imp)
		}
		pretty.Println("imps:", imps)
	}

	panic("bar")

	return file, nil
}

// ref: https://msdn.microsoft.com/en-us/library/ms809762.aspx

// An importDesc is an import descriptor.
type importDesc struct {
	// Import name table RVA.
	ImportNameTableRVA uint32
	// Time stamp.
	Date uint32
	// Forward chain; index into importAddressTableRVA for forwarding a function
	// to another DLL.
	ForwardChain uint32
	// DLL name RVA.
	DLLNameRVA uint32
	// Import address table RVA.
	ImportAddressTableRVA uint32
}

// An importName specifies the name of an import.
type importName struct {
	// Approximate ordinal number (used by loader to initiate binary search).
	Ordinal uint16
	// Name of the import.
	Name string
}

// parsePerm returns the memory access permissions represented by the given PE
// image characteristics.
func parsePerm(char uint32) bin.Perm {
	// Characteristics.
	const (
		// permR specifies that the memory is readable.
		permR = 0x40000000
		// permW specifies that the memory is writeable.
		permW = 0x80000000
		// permX specifies that the memory is executable.
		permX = 0x20000000
	)
	var perm bin.Perm
	if char&permR != 0 {
		perm |= bin.PermR
	}
	if char&permW != 0 {
		perm |= bin.PermW
	}
	if char&permX != 0 {
		perm |= bin.PermX
	}
	return perm
}
