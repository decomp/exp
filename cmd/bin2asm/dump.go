package main

import (
	"bufio"
	"debug/pe"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// dumpMainAsm dumps the main.asm file of the executable.
func dumpMainAsm(file *pe.File) error {
	t, err := parseTemplate("main.asm.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	var sectNames []string
	for _, sect := range file.Sections {
		sectName := strings.Replace(sect.Name, ".", "_", -1)
		sectNames = append(sectNames, sectName)
	}
	// Store output.
	if err := writeFile(t, "main.asm", sectNames); err != nil {
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
	if err := t.Execute(bw, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// parseTemplate parses and returns the given template.
func parseTemplate(filename string) (*template.Template, error) {
	path := filepath.Join(bin2asmDir, filename)
	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}
