// The dump_imports tool dumps the imports of a PE binary in NASM syntax.
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/kr/pretty"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewrev/pe"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		if err := dumpImports(path); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// dumpImports dumps the imports of the given executable in NASM syntax.
func dumpImports(path string) error {
	// Parse PE file.
	file, err := pe.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	optHdr, err := file.OptHeader()
	if err != nil {
		return errors.WithStack(err)
	}
	sectHdrs, err := file.SectHeaders()
	if err != nil {
		return errors.WithStack(err)
	}
	// data returns the data at the given address.
	data := func(relAddr uint32) []byte {
		for _, sectHdr := range sectHdrs {
			start := sectHdr.RelAddr
			end := start + sectHdr.Size
			if start <= relAddr && relAddr < end {
				data, err := file.Section(sectHdr)
				if err != nil {
					panic(fmt.Errorf("unable to read section contents at RVA 0x%08X; %v", relAddr, err))
				}
				offset := relAddr - start
				return data[offset:]
			}
		}
		panic(fmt.Errorf("unable to locate section containing RVA 0x%08X", relAddr))
	}
	const (
		importTableIndex        = 1
		importAddressTableIndex = 12
	)
	itDir := optHdr.DataDirs[importTableIndex]
	iatDir := optHdr.DataDirs[importAddressTableIndex]

	it := data(itDir.RelAddr)
	it = it[:itDir.Size]
	fmt.Println("import table:")
	fmt.Println(hex.Dump(it))

	iat := data(iatDir.RelAddr)
	iat = iat[:iatDir.Size]
	fmt.Println("import address table:")
	fmt.Println(hex.Dump(iat))

	// Parse imports.
	if err := parseImports(it, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

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

// parseImports parses the given import table and import address table.
func parseImports(it []byte, data func(addr uint32) []byte) error {
	buf := &bytes.Buffer{}
	bufINT := &bytes.Buffer{}
	var impDesc importDesc
	r := bytes.NewReader(it)
	const itHeader = `
; === [ import table ] =========================================================
;
;     Array of IMAGE_IMPORT_DESCRIPTOR structures, which is terminated by an
;     empty struct.
;
; ------------------------------------------------------------------------------

import_table:
`
	buf.WriteString(itHeader[1:])

	const intHeader = `
; === [ Import Name Tables (INTs) ] ============================================
;
;     Each Import Name Table (INT) consists of an array of IMAGE_THUNK_DATA
;     structures, which are terminated by an empty struct.
;
;     This entire table is UNUSED. It is sometimes refered to as the
;     "hint table", which is never overwritten or altered.
; ------------------------------------------------------------------------------

`
	bufINT.WriteString(intHeader[1:])

	const iatHeader = `
; === [ Import Address Tables (IATs) ] =========================================
;
;     Each Import Address Table (IAT) consists of an array of IMAGE_THUNK_DATA
;     structures, which are terminated by an empty struct.
;
;     The linker will overwrite these DWORDs with the actuall address of the
;     imported functions
;
; ------------------------------------------------------------------------------

iat:

`
	bufINT.WriteString(iatHeader[1:])
	var dlls []*DLL
	for {
		if err := binary.Read(r, binary.LittleEndian, &impDesc); err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return errors.WithStack(err)
		}
		if impDesc == (importDesc{}) {
			buf.WriteString("   times 5              dd      0x00000000\n\n")
			break
		}
		fmt.Println("import descriptor:", impDesc)
		dllName := parseString(data(impDesc.DLLNameRVA))

		const importDescFormat = `
                        dd      int_%s - IMAGE_BASE      ; pINT_first_trunk
                        dd      0x%08X                         ; TimeDateStamp
                        dd      0x%08X                         ; pForwardChain
                        dd      sz%s_dll - IMAGE_BASE
                        dd      iat_%s - IMAGE_BASE      ; pIAT_first_trunk

`
		name := strings.ToLower(pathutil.TrimExt(dllName))
		fmt.Fprintf(buf, importDescFormat[1:], name, impDesc.Date, impDesc.ForwardChain, strings.Title(name), name)

		fmt.Println("  => DLL name:", dllName)
		dllINT := data(impDesc.ImportNameTableRVA)
		r := bytes.NewReader(dllINT)
		if err := dumpINT(name, r, bufINT, data); err != nil {
			return errors.WithStack(err)
		}

		dllIAT := data(impDesc.ImportAddressTableRVA)
		r = bytes.NewReader(dllIAT)
		if err := dumpIAT(name, r, bufINT, data); err != nil {
			return errors.WithStack(err)
		}
		r = bytes.NewReader(dllIAT)
		dll, err := parseDLL(dllName, r, data)
		if err != nil {
			return errors.WithStack(err)
		}
		dll.Addr = impDesc.DLLNameRVA
		dlls = append(dlls, dll)
	}
	const importTableFooter = `
   import_table_size    equ     $ - import_table

; === [/ import table ] ========================================================

`
	buf.WriteString(importTableFooter[1:])

	const intHeaderFooter = `
; === [/ Import Name Tables (INTs) ] ===========================================

`
	bufINT.WriteString(intHeaderFooter[1:])

	const iatHeaderFooter = `
   iat_size             equ     $ - iat

; === [/ Import Address Tables (IATs) ] ========================================

`
	bufINT.WriteString(iatHeaderFooter[1:])

	if _, err := buf.Write(bufINT.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	if _, err := buf.Write(bufINT.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	less := func(i, j int) bool {
		return dlls[i].Addr < dlls[j].Addr
	}
	sort.Slice(dlls, less)
	pretty.Println(dlls)
	dumpDLLs(dlls, buf)

	if err := os.MkdirAll("_dump_", 0755); err != nil {
		return errors.WithStack(err)
	}
	if err := ioutil.WriteFile("_dump_/_idata.asm", buf.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func dumpDLLs(dlls []*DLL, w io.Writer) {

	const dllsHeader = `
; === [ dll and function names ] ===============================================
;
;    each dll stores an array of IMAGE_IMPORT_BY_NAME structures and a string
;    corresponding to it's dll name.
;
; ------------------------------------------------------------------------------

`
	fmt.Fprint(w, dllsHeader[1:])
	for _, dll := range dlls {
		const dllHeaderFormat = `
; --- [ %s ] ---------------------------------------------------------

`
		fmt.Fprintf(w, dllHeaderFormat[1:], dll.Name)
		for _, f := range dll.Funcs {
			const funcFormat = `
imp_%s:
                        dw      0x%04X
                        db      '%s', 0x00 ; 0x%08X
                        align 2, db 0x00

`
			fmt.Fprintf(w, funcFormat[1:], f.Name, f.Ordinal, f.Name, f.Addr)
		}
		const dllFooterFormat = `
sz%s:
                        db      '%s', 0x00 ; 0x%08X
                        align 2, db 0x00

`
		fmt.Fprintf(w, dllFooterFormat[1:], strings.Title(strings.ToLower(strings.Replace(dll.Name, ".", "_", -1))), dll.Name, dll.Addr)
	}
	const dllsFooter = `
; === [/ dll and function names ] ==============================================
`
	fmt.Fprint(w, dllsFooter[1:])
}

func dumpINT(dllName string, r io.Reader, w io.Writer, data func(addr uint32) []byte) error {
	const intHeaderFormat = `
; --- [ %s.dll ] ---------------------------------------------------------

int_%s:
`
	fmt.Fprintf(w, intHeaderFormat[1:], dllName, dllName)
	for {
		var addr uint32
		if err := binary.Read(r, binary.LittleEndian, &addr); err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return errors.WithStack(err)
		}
		if addr == 0 {
			const intNullEntry = `
                        dd      0x00000000
`
			fmt.Fprintln(w, intNullEntry[1:])
			break
		}
		fmt.Printf("addr: 0x%08X\n", addr)
		if addr&0x80000000 != 0 {
			// addr is an encoded ordinal, not an address.
			ordinal := addr &^ 0x80000000
			fmt.Println("ordinal:", ordinal)
			const intOrdinalFormat = `
                        dd      0x80000000 | %d
`
			fmt.Fprintf(w, intOrdinalFormat[1:], ordinal)
		} else {
			var ordinal uint16
			if err := binary.Read(bytes.NewReader(data(addr)), binary.LittleEndian, &ordinal); err != nil {
				return errors.WithStack(err)
			}
			funcName := parseString(data(addr + 2))
			fmt.Printf("function: %s (%d)\n", funcName, ordinal)
			const intEntryFormat = `
                        dd      imp_%s - IMAGE_BASE
`
			fmt.Fprintf(w, intEntryFormat[1:], funcName)
		}
	}
	return nil
}

type DLL struct {
	Name  string
	Addr  uint32
	Funcs []*Func
}

type Func struct {
	Addr    uint32
	Ordinal uint16
	Name    string
}

func parseDLL(dllName string, r io.Reader, data func(addr uint32) []byte) (*DLL, error) {
	dll := &DLL{
		Name: dllName,
	}
	for {
		var addr uint32
		if err := binary.Read(r, binary.LittleEndian, &addr); err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return nil, errors.WithStack(err)
		}
		if addr == 0 {
			break
		}
		if addr&0x80000000 != 0 {
			// ordinal, nothing to do.
			continue
		}
		f := &Func{
			Addr: addr,
		}
		dll.Funcs = append(dll.Funcs, f)
		if err := binary.Read(bytes.NewReader(data(addr)), binary.LittleEndian, &f.Ordinal); err != nil {
			return nil, errors.WithStack(err)
		}
		f.Name = parseString(data(addr + 2))
	}
	less := func(i, j int) bool {
		return dll.Funcs[i].Addr < dll.Funcs[j].Addr
	}
	sort.Slice(dll.Funcs, less)
	return dll, nil
}

func dumpIAT(dllName string, r io.Reader, w io.Writer, data func(addr uint32) []byte) error {
	const iatHeaderFormat = `
; --- [ %s.dll ] ---------------------------------------------------------

iat_%s:
`
	fmt.Fprintf(w, iatHeaderFormat[1:], dllName, dllName)
	for {
		var addr uint32
		if err := binary.Read(r, binary.LittleEndian, &addr); err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return errors.WithStack(err)
		}
		if addr == 0 {
			const iatNullEntry = `
                        dd      0x00000000
`
			fmt.Fprintln(w, iatNullEntry[1:])
			break
		}
		fmt.Printf("addr: 0x%08X\n", addr)
		if addr&0x80000000 != 0 {
			// addr is an encoded ordinal, not an address.
			ordinal := addr &^ 0x80000000
			fmt.Println("ordinal:", ordinal)
			const iatOrdinalFormat = `
  ia_%s_%d:
                        dd      0x80000000 | %d
`
			fmt.Fprintf(w, iatOrdinalFormat[1:], dllName, ordinal, ordinal)
		} else {
			var ordinal uint16
			if err := binary.Read(bytes.NewReader(data(addr)), binary.LittleEndian, &ordinal); err != nil {
				return errors.WithStack(err)
			}
			funcName := parseString(data(addr + 2))
			fmt.Printf("function: %s (%d)\n", funcName, ordinal)
			const iatEntryFormat = `
  ia_%s:
                        dd      imp_%s - IMAGE_BASE
`
			fmt.Fprintf(w, iatEntryFormat[1:], funcName, funcName)
		}
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// parseString converts the given NULL-terminated string to a Go string.
func parseString(buf []byte) string {
	pos := bytes.IndexByte(buf, '\x00')
	if pos == -1 {
		panic(fmt.Errorf("unable to locate NULL-terminated string in % 02X", buf))
	}
	return string(buf[:pos])
}
