package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// dumpOverlay dumps the overlay of the PE file in NASM syntax.
func dumpOverlay(overlay []byte) error {
	buf := &bytes.Buffer{}
	for _, b := range overlay {
		// Dump data.
		//
		//    db      0x44 ; 'D'
		char := ""
		if isPrint(b) {
			char = fmt.Sprintf(" ; %q", b)
		}
		fmt.Fprintf(buf, "        db      0x%02X%s\n", b, char)
	}
	filename := "overlay.asm"
	outPath := filepath.Join(outDir, filename)
	dbg.Printf("creating %q\n", outPath)
	if err := ioutil.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
