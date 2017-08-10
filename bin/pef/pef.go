// Package pef provides access to PEF (Preferred Executable Format) files.
package pef

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

// Register PEF format.
func init() {
	// Preferred Executable Format (PEF) format.
	//
	//    4A 6F 79 21 70 65 66 66  |Joy!peff|
	const magic = "Joy!peff"
	bin.RegisterFormat("pef", magic, Parse)
}

// ParseFile parses the given PEF binary executable, reading from path.
func ParseFile(path string) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given PEF binary executable, reading from r.
//
// Users are responsible for closing r.
func Parse(r io.ReaderAt) (*bin.File, error) {
	// Open PEF file.
	f, err := NewFile(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse machine architecture.
	file := &bin.File{}
	for _, container := range f.Containers {
		var arch bin.Arch
		switch container.Architecture {
		case "pwpc":
			arch = bin.ArchPowerPC_32
		default:
			panic(fmt.Errorf("support for machine architecture %q not yet implemented", container.Architecture))
		}
		if file.Arch != 0 && arch != file.Arch {
			panic(fmt.Errorf("support for multiple machine architectures not yet implemented; prev %q, new %q", file.Arch, arch))
		}
		file.Arch = arch
	}

	// Parse sections.
	for _, container := range f.Containers {
		for _, s := range container.Sections {
			data, err := s.Data()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			perm := parsePerm(s.SectionKind)
			offset := container.Offset + uint64(s.ContainerOffset)
			sect := &bin.Section{
				Addr:     bin.Address(s.DefaultAddress),
				Offset:   offset,
				Data:     data,
				FileSize: int(s.PackedSize),
				MemSize:  int(s.TotalSize),
				Perm:     perm,
			}
			file.Sections = append(file.Sections, sect)
		}
	}

	return file, nil
}

// NewFile creates a new File for accessing a PEF binary in an underlying
// reader.
//
// Users are responsible for closing r.
func NewFile(r io.ReaderAt) (*File, error) {
	f, err := parseFile(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return f, nil
}

// A File is PEF file.
type File struct {
	// PEF containers.
	Containers []*Container
}

// parseFile parses and returns a PEF file.
func parseFile(r io.ReaderAt) (*File, error) {
	var offset int64
	f := &File{}
	for {
		sr := io.NewSectionReader(r, offset, math.MaxInt64)
		container, n, err := parseContainer(sr)
		if err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return nil, errors.WithStack(err)
		}
		offset += n
		f.Containers = append(f.Containers, container)
	}
	return f, nil
}

// A Container is a PEF container.
type Container struct {
	// PEF container header.
	*ContainerHeader
	// File offset of the container.
	Offset uint64
	// PEF sections.
	Sections []*Section
}

// parseContainer parses and returns a PEF container.
func parseContainer(r io.ReaderAt) (*Container, int64, error) {
	// Overview of the structure of a PEF container.
	//
	//    Container header
	//    Section headers; zero or more
	//    Section name table
	//    Section contents; zero or more

	// Parse PEF container header.
	var offset int64
	hdr, n, err := parseContainerHeader(r)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	offset += n
	container := &Container{
		ContainerHeader: hdr,
	}

	// Parse section headers.
	for i := uint16(0); i < hdr.SectionCount; i++ {
		sr := io.NewSectionReader(r, offset, math.MaxInt64)
		sectHdr, n, err := parseSectionHeader(sr)
		if err != nil {
			return nil, 0, errors.WithStack(err)
		}
		offset += n
		sect := &Section{
			SectionHeader: sectHdr,
			ReaderAt:      io.NewSectionReader(r, int64(sectHdr.ContainerOffset), int64(sectHdr.PackedSize)),
		}
		container.Sections = append(container.Sections, sect)
	}

	// The PEF container section name table contains the names of the sections
	// stored as C-style null-terminated character strings. The strings have no
	// specified alignment. Note that the section name table must immediately
	// follow the section headers in the container.

	// TODO: Parse string table.

	// Parse Loader section.
	for _, sect := range container.Sections {
		fmt.Println("kind:", sect.SectionKind)
		if sect.SectionKind == kindLoader {
			if err := parseLoaderSection(sect); err != nil {
				return nil, 0, errors.WithStack(err)
			}
		}
	}

	// Calculate offset based on the end of the last section.
	for _, sect := range container.Sections {
		x := int64(sect.SectionHeader.ContainerOffset + sect.SectionHeader.PackedSize)
		if offset < x {
			offset = x
		}
	}
	// Adjust offset based on alignment.
	//
	// When the container is not file-mapped, the overall container alignment is
	// 16 bytes.
	//
	// ref: https://web.archive.org/web/20020111211702/http://developer.apple.com:80/techpubs/mac/runtimehtml/RTArch-92.html
	if rem := offset % 16; rem != 0 {
		offset += 16 - rem
	}
	return container, offset, nil
}

// A ContainerHeader represents a single PEF container header.
//
// ref: https://web.archive.org/web/20020219190852/http://developer.apple.com:80/techpubs/mac/runtimehtml/RTArch-91.html
type ContainerHeader struct {
	// Magic header: "Joy!"
	Tag1 string
	// Magic header: "peff"
	Tag2 string
	// Machine architecture.
	//    "pwpc" for PowerPC
	//    "m68k" for Motorola 68K
	Architecture string
	// PEF container format version.
	FormatVersion uint32
	// PEF container creation date.
	DateTimeStamp time.Time
	// Old definition version.
	OldDefVersion uint32
	// Old implementation version.
	OldImpVersion uint32
	// Current version.
	CurrentVersion uint32
	// Number of sections.
	SectionCount uint16
	// Number of instantiated sections.
	InstSectionCount uint16
}

// parseContainerHeader parses and returns a PEF container header.
func parseContainerHeader(r io.ReaderAt) (*ContainerHeader, int64, error) {
	// PEF container header.
	//
	// ref: https://web.archive.org/web/20020219190852/http://developer.apple.com:80/techpubs/mac/runtimehtml/RTArch-91.html
	const containerHeaderSize = 40
	type containerHeader struct {
		Tag1          [4]byte // 4 bytes; "Joy!"
		Tag2          [4]byte // 4 bytes; "peff"
		Architecture  [4]byte // 4 bytes; "pwpc" or "m68k"
		FormatVersion uint32
		// In Macintosh time-measurement scheme (number of seconds measured from
		// January 1, 1904).
		DateTimeStamp    uint32
		OldDefVersion    uint32
		OldImpVersion    uint32
		CurrentVersion   uint32
		SectionCount     uint16
		InstSectionCount uint16
		_                uint32 // reserved
	}
	buf := make([]byte, containerHeaderSize)
	if _, err := r.ReadAt(buf, 0); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	v := &containerHeader{}
	if err := binary.Read(bytes.NewReader(buf), binary.BigEndian, v); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	epoch := time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Duration(v.DateTimeStamp) * time.Second
	date := epoch.Add(dur)
	hdr := &ContainerHeader{
		Tag1:             string(v.Tag1[:]),
		Tag2:             string(v.Tag2[:]),
		Architecture:     string(v.Architecture[:]),
		FormatVersion:    v.FormatVersion,
		DateTimeStamp:    date,
		OldDefVersion:    v.OldDefVersion,
		OldImpVersion:    v.OldImpVersion,
		CurrentVersion:   v.CurrentVersion,
		SectionCount:     v.SectionCount,
		InstSectionCount: v.InstSectionCount,
	}
	return hdr, containerHeaderSize, nil
}

// A Section is a PEF section.
type Section struct {
	// PEF section header.
	*SectionHeader
	io.ReaderAt
}

// Data reads and returns the contents of the PEF section.
func (sect *Section) Data() ([]byte, error) {
	buf := make([]byte, sect.PackedSize)
	if _, err := sect.ReadAt(buf, 0); err != nil {
		return nil, errors.WithStack(err)
	}
	return buf, nil
}

// A SectionHeader is a PEF section header.
type SectionHeader struct {
	// Offset from start of section name table to section name; or -1 if
	// section has no name.
	NameOffset int32
	// Preferred address at which to place the section's instance.
	DefaultAddress uint32
	// Size in bytes required by the section's contents at execution time. For a
	// code section, this size is merely the size of the executable code. For a
	// data section, this size indicates the sum of the size of the initialized
	// data plus the size of any zero-initialized data. Zero-initialized data
	// appears at the end of a section's contents and its length is exactly the
	// difference of the TotalSize and UnpackedSize values.
	//
	// For noninstantiated sections, this field is ignored.
	TotalSize uint32
	// Size of the section's contents that is explicitly initialized from the
	// container. For code sections, this field is the size of the executable
	// code. For an unpacked data section, this field indicates only the size of
	// the initialized data. For packed data this is the size to which the
	// compressed contents expand. The UnpackedSize value also defines the
	// boundary between the explicitly initialized portion and the zero-
	// initialized portion.
	//
	// For noninstantiated sections, this field is ignored.
	UnpackedSize uint32
	// Size in bytes of a section's contents in the container. For code sections,
	// this field is the size of the executable code. For an unpacked data
	// section, this field indicates only the size of the initialized data. For a
	// packed data section this field is the size of the pattern description
	// contained in the section.
	PackedSize uint32
	// Offset from the beginning of the container to the start of the section's
	// contents.
	ContainerOffset uint32
	// Indicates the type of section as well as any special attributes.
	SectionKind uint8
	// Controls how the section information is shared among processes.
	ShareKind uint8
	// Indicates the desired alignment for instantiated sections in memory as a
	// power of 2.
	Alignment uint8
}

// parseSectionHeader parses and returns a PEF section header.
func parseSectionHeader(r io.ReaderAt) (*SectionHeader, int64, error) {
	// PEF section header.
	//
	// ref: https://web.archive.org/web/20020111211702/http://developer.apple.com:80/techpubs/mac/runtimehtml/RTArch-92.html
	const sectionHeaderSize = 28
	buf := make([]byte, sectionHeaderSize)
	if _, err := r.ReadAt(buf, 0); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	hdr := &SectionHeader{}
	if err := binary.Read(bytes.NewReader(buf), binary.BigEndian, hdr); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	if hdr.NameOffset != -1 {
		panic("support for section name table not yet implemented")
	}
	return hdr, sectionHeaderSize, nil
}

// Section kinds.
const (
	// Read-only executable code.
	kindCode = 0
	// Read/write data.
	kindUnpackedData = 1
	// Read/write data.
	kindPatternInitializedData = 2
	// Read-only data.
	kindConstant = 3
	// Contains information about imports, exports, and entry points.
	kindLoader = 4
	// Reserved for future use.
	kindDebug = 5
	// Read/write, executable code.
	kindExecutableData = 6
	// Reserved for future use.
	kindException = 7
	// Reserved for future use.
	kindTraceback = 8
)

// parsePerm returns the memory access permissions represented by the given PEF
// section kind.
func parsePerm(kind uint8) bin.Perm {
	var perm bin.Perm
	switch kind {
	case kindCode, kindUnpackedData, kindPatternInitializedData, kindConstant, kindExecutableData:
		perm |= bin.PermR
	}
	switch kind {
	case kindUnpackedData, kindPatternInitializedData, kindExecutableData:
		perm |= bin.PermW
	}
	switch kind {
	case kindCode, kindExecutableData:
		perm |= bin.PermX
	}
	return perm
}

// parseLoaderSection parses the given Loader section.
func parseLoaderSection(sect *Section) error {
	// Overview of the structure of a PEF Loader section.
	//
	//    Loader header
	//    Imported library table
	//    Imported symbol table
	//    Relocation headers table
	//    Relocations
	//    Loader string table
	//    Export hash table
	//    Export key table
	//    Exported symbol table

	const loaderHeaderSize = 56
	type LoaderHeader struct {
		MainSection              int32
		MainOffset               uint32
		InitSection              int32
		InitOffset               uint32
		TermSection              int32
		TermOffset               uint32
		ImportedLibraryCount     uint32
		TotalImportedSymbolCount uint32
		RelocSectionCount        uint32
		RelocInstrOffset         uint32
		LoaderStringsOffset      uint32
		ExportHashOffset         uint32
		ExportHashTablePower     uint32
		ExportedSymbolCount      uint32
	}
	buf := make([]byte, loaderHeaderSize)
	if _, err := sect.ReadAt(buf, 0); err != nil {
		return errors.WithStack(err)
	}
	hdr := &LoaderHeader{}
	if err := binary.Read(bytes.NewReader(buf), binary.BigEndian, hdr); err != nil {
		return errors.WithStack(err)
	}

	pretty.Println(hdr)

	return nil
}
