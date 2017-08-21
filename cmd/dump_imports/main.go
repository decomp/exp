// The imports tool dumps the imports of a PE binary in NASM syntax.
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
	"strings"

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
		dllIAT := data(impDesc.ImportAddressTableRVA)
		r := bytes.NewReader(dllIAT)
		if err := dumpINT(name, r, bufINT, data); err != nil {
			return errors.WithStack(err)
		}
	}
	const importTableFooter = `
   import_table_size    equ     $ - import_table

; === [/ import table ] ========================================================

`
	buf.WriteString(importTableFooter[1:])

	if _, err := buf.Write(bufINT.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	if err := os.MkdirAll("_dump_", 0755); err != nil {
		return errors.WithStack(err)
	}
	if err := ioutil.WriteFile("_dump_/_idata.asm", buf.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
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

// ### [ Helper functions ] ####################################################

// parseString converts the given NULL-terminated string to a Go string.
func parseString(buf []byte) string {
	pos := bytes.IndexByte(buf, '\x00')
	if pos == -1 {
		panic(fmt.Errorf("unable to locate NULL-terminated string in % 02X", buf))
	}
	return string(buf[:pos])
}
