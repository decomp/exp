// Package pe provides access to PE binary executables.
package pe

import (
	"debug/pe"
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
)

// Parse parses the given PE binary executable.
func Parse(path string) (*bin.File, error) {
	// Open PE file.
	f, err := pe.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()

	// Determine machine architecture.
	file := &bin.File{}
	switch f.FileHeader.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		file.Arch = bin.ArchX86_32
	case pe.IMAGE_FILE_MACHINE_AMD64:
		file.Arch = bin.ArchX86_64
	default:
		panic(fmt.Errorf("support for machine architecture %v not yet implemented", f.FileHeader.Machine))
	}

	// Parse optional header.
	var (
		imageBase uint64
		//idataBase uint64
		//idataSize uint64
	)
	switch opt := f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		file.Entry = bin.Address(opt.AddressOfEntryPoint + opt.ImageBase)
		imageBase = uint64(opt.ImageBase)
		//idataBase = uint64(opt.DataDirectory[12].VirtualAddress)
		//idataSize = uint64(opt.DataDirectory[12].Size)
	case *pe.OptionalHeader64:
		file.Entry = bin.Address(opt.AddressOfEntryPoint) + bin.Address(opt.ImageBase)
		imageBase = uint64(opt.ImageBase)
		//idataBase = uint64(opt.DataDirectory[12].VirtualAddress)
		//idataSize = uint64(opt.DataDirectory[12].Size)
	default:
		panic(fmt.Errorf("support for optional header type %T not yet implemented", opt))
	}

	// Parse sections.
	for _, s := range f.Sections {
		data, err := s.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// Parse access permissions.
		perm := parsePerm(s.Characteristics)
		sect := &bin.Section{
			Name: s.Name,
			Addr: bin.Address(s.VirtualAddress) + bin.Address(imageBase),
			Data: data,
			Perm: perm,
		}
		file.Sections = append(file.Sections, sect)
	}

	// TODO: Parse imports.

	return file, nil
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
