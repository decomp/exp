package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mewkiz/pkg/goutil"
	"github.com/mewrev/pe"
	"github.com/pkg/errors"
)

// bin2asmDir specifies the source directory of the bin2asm command.
var bin2asmDir string

func init() {
	var err error
	bin2asmDir, err = goutil.SrcDir("github.com/decomp/exp/cmd/bin2asm")
	if err != nil {
		panic(fmt.Errorf("unable to locate source directory of bin2asm; %v", err))
	}
}

// dumpCommon dumps a common include file of the executable.
func dumpCommon(file *pe.File) error {
	buf := &bytes.Buffer{}
	const commonFormat = `
%%ifndef __COMMON_INC__
%%define __COMMON_INC__

%%define IMAGE_BASE      0x%08X
%%define CODE_BASE       0x%08X
%%define DATA_BASE       0x%08X

   hdr_vstart           equ     IMAGE_BASE
`
	optHdr, err := file.OptHeader()
	if err != nil {
		return errors.WithStack(err)
	}
	sectHdrs, err := file.SectHeaders()
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Fprintf(buf, commonFormat[1:], optHdr.ImageBase, optHdr.CodeBase, optHdr.DataBase)

	for _, sectHdr := range sectHdrs {
		rawName := sectHdr.Name
		sectName := strings.Replace(rawName, ".", "_", -1)
		fmt.Fprintf(buf, "   %s_vstart         equ     IMAGE_BASE + 0x%08X\n", sectName, sectHdr.RelAddr)
	}
	buf.WriteString("\n")
	for _, sectHdr := range sectHdrs {
		if sectHdr.VirtSize > sectHdr.Size {
			// Section with uninitialized data.
			buf.WriteString(";")
		}
		rawName := sectHdr.Name
		sectName := strings.Replace(rawName, ".", "_", -1)
		fmt.Fprintf(buf, "   %s_size           equ     0x%08X\n", sectName, sectHdr.Size)
	}
	buf.WriteString("\n")
	for _, sectHdr := range sectHdrs {
		if sectHdr.VirtSize <= sectHdr.Size {
			// Section with padding.
			buf.WriteString(";")
		}
		rawName := sectHdr.Name
		sectName := strings.Replace(rawName, ".", "_", -1)
		fmt.Fprintf(buf, "   %s_vsize          equ     0x%08X\n", sectName, sectHdr.VirtSize)
	}
	buf.WriteString("\n")

	const bitsHeader = `
BITS 32

SECTION hdr    vstart=hdr_vstart
`
	buf.WriteString(bitsHeader[1:])

	prev := "hdr"
	for _, sectHdr := range sectHdrs {
		rawName := sectHdr.Name
		sectName := strings.Replace(rawName, ".", "_", -1)
		fmt.Fprintf(buf, "SECTION %s  vstart=%s_vstart  follows=%s\n", rawName, sectName, prev)
		prev = rawName
	}
	buf.WriteString("\n")
	buf.WriteString("%endif ; %ifndef __COMMON_INC__\n")

	// Store output.
	outPath := filepath.Join(outDir, "common.inc")
	dbg.Printf("creating %q\n", outPath)
	if err := ioutil.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
