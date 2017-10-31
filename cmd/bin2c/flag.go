package main

import (
	"fmt"
	"go/ast"
	"log"

	"github.com/mewkiz/pkg/errutil"
)

type Flag uint8

func (flag Flag) String() string {
	m := map[Flag]string{
		CF: "cf",
		OF: "of",
		SF: "sf",
		ZF: "zf",
	}
	if s, ok := m[flag]; ok {
		return s
	}
	return fmt.Sprintf("<unknown flag %d>", uint8(flag))
}

const (
	CF Flag = iota + 1
	OF
	SF
	ZF
)

// flags maps flag names to their corresponding Go identifiers.
var flags = map[Flag]*ast.Ident{
	CF: ast.NewIdent("cf"),
	OF: ast.NewIdent("of"),
	SF: ast.NewIdent("sf"),
	ZF: ast.NewIdent("zf"),
}

// getFlag converts flag into a corresponding Go expression.
func getFlag(flag Flag) ast.Expr {
	if expr, ok := flags[flag]; ok {
		return expr
	}
	log.Fatal(errutil.Newf("unable to lookup identifer for flag %q", flag))
	panic("unreachable")
}
