package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode"

	"github.com/mewrev/pe"
	"github.com/pkg/errors"
)

// dumpMainAsm dumps the main.asm file of the executable.
func dumpMainAsm(file *pe.File, hasOverlay bool) error {
	t, err := parseTemplate("main.asm.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	sectHdrs, err := file.SectHeaders()
	if err != nil {
		return errors.WithStack(err)
	}
	var data []string
	for _, sectHdr := range sectHdrs {
		sectName := underline(sectHdr.Name)
		data = append(data, sectName)
	}
	if hasOverlay {
		data = append(data, "overlay")
	}
	// Store output.
	if err := writeFile(t, "main.asm", data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// dumpPEHeaderAsm dumps the pe-hdr.asm file of the executable.
func dumpPEHeaderAsm(file *pe.File) error {
	t, err := parseTemplate("pe-hdr.asm.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	optHdr, err := file.OptHeader()
	if err != nil {
		return errors.WithStack(err)
	}
	dosHdr, err := file.DOSHeader()
	if err != nil {
		return errors.WithStack(err)
	}
	dosStub, err := file.DOSStub()
	if err != nil {
		return errors.WithStack(err)
	}
	fileHdr, err := file.FileHeader()
	if err != nil {
		return errors.WithStack(err)
	}
	sectAlignKB := optHdr.SectAlign / 1024
	sectHdrs, err := file.SectHeaders()
	if err != nil {
		return errors.WithStack(err)
	}
	// Calculate size of sections containing data.
	var dataSizes []string
	for _, sectHdr := range sectHdrs {
		if sectHdr.Flags&pe.SectFlagData == 0 {
			// Ignore non-data sections.
			continue
		}
		if sectHdr.Size < sectHdr.VirtSize {
			// Section contains uninitialized data.
			dataSize := fmt.Sprintf("%s_vsize", underline(sectHdr.Name))
			dataSizes = append(dataSizes, dataSize)
		} else {
			// Section is padded.
			dataSize := fmt.Sprintf("%s_size", underline(sectHdr.Name))
			dataSizes = append(dataSizes, dataSize)
		}
	}
	// Store output.
	data := map[string]interface{}{
		"OptHdr":      optHdr,
		"DosHdr":      dosHdr,
		"DOSStub":     dosStub,
		"FileHdr":     fileHdr,
		"SectAlignKB": sectAlignKB,
		"SectHdrs":    sectHdrs,
		"DataSizes":   strings.Join(dataSizes, " + "),
		"DataDirs":    optHdr.DataDirs,
	}
	if err := writeFile(t, "pe-hdr.asm", data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// writeFile applies a parsed template to the specified data object, writing the
// output to the specified file.
func writeFile(t *template.Template, filename string, data interface{}) error {
	path := filepath.Join(outDir, filename)
	dbg.Printf("creating %q\n", path)
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	defer bw.Flush()
	tw := tabwriter.NewWriter(bw, 1, 8, 1, ' ', 0)
	if err := t.Execute(tw, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// parseTemplate parses and returns the given template.
func parseTemplate(filename string) (*template.Template, error) {
	funcMap := map[string]interface{}{
		"isprint":   isPrint,
		"underline": underline,
		"nameArray": nameArray,
		"ui16":      ui16,
		"ui32":      ui32,
	}
	path := filepath.Join(bin2asmDir, filename)
	t, err := template.New(filename).Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

// isPrint reports if the given byte is printable.
func isPrint(b byte) bool {
	if b >= 0x7F {
		return false
	}
	return unicode.IsPrint(rune(b))
}

// underline replaces dot characters in the given string with underscore
// characters.
func underline(s string) string {
	return strings.Replace(s, ".", "_", -1)
}

// nameArray converts the given string to a byte array of 8 characters, pretty-
// printed in NASM syntax.
func nameArray(s string) string {
	out := fmt.Sprintf("%q", s)
	for i := len(s); i < 8; i++ {
		out += ", 0x00"
	}
	return out
}

// ui16 converts x to a 16-bit unsigned integer.
func ui16(x interface{}) uint16 {
	switch x := x.(type) {
	case pe.Arch:
		return uint16(x)
	case pe.DLLFlag:
		return uint16(x)
	case pe.Flag:
		return uint16(x)
	case pe.OptState:
		return uint16(x)
	case pe.Subsystem:
		return uint16(x)
	default:
		panic(fmt.Errorf("main.ui16: support for type %T not yet implemented", x))
	}
}

// ui32 converts x to a 32-bit unsigned integer.
func ui32(x interface{}) uint32 {
	switch x := x.(type) {
	case pe.SectFlag:
		return uint32(x)
	case pe.Time:
		return uint32(x)
	default:
		panic(fmt.Errorf("main.ui32: support for type %T not yet implemented", x))
	}
}
