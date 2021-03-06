// The dot2png tool converts DOT files to PNG images.
//
// Usage:
//
//    dot2png FILE.dot...
//
// Flags:
//
//    -f    force overwrite existing images
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewkiz/pkg/pathutil"
)

func usage() {
	const use = `
Usage: dot2png FILE.dot...
Convert DOT files to PNG images.

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line options.
	var (
		// force specifies whether to force overwrite existing images.
		force bool
	)
	flag.BoolVar(&force, "f", false, "force overwrite existing images")
	flag.Usage = usage
	flag.Parse()

	// Convert DOT files to PNG images.
	for _, dotPath := range flag.Args() {
		if err := convert(dotPath, force); err != nil {
			log.Fatal(err)
		}
	}
}

// convert converts the provided DOT file to a PNG image.
func convert(dotPath string, force bool) error {
	pngPath := pathutil.TrimExt(dotPath) + ".png"

	// Skip existing files unless the "-f" flag is set.
	if !force {
		dotStat, err := os.Stat(dotPath)
		if err != nil {
			return errutil.Err(err)
		}
		pngStat, err := os.Stat(pngPath)
		if err != nil {
			if !os.IsNotExist(err) {
				return errutil.Err(err)
			}
		} else {
			dotMod, pngMod := dotStat.ModTime(), pngStat.ModTime()
			if dotMod.Before(pngMod) {
				// PNG file is newer than DOT file, ignore.
				return nil
			}
		}
	}

	// Convert the DOT file to a PNG image.
	cmd := exec.Command("dot", "-Tpng", "-o", pngPath, dotPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Creating: %q", pngPath)
	if err := cmd.Run(); err != nil {
		return errutil.Err(err)
	}

	return nil
}
