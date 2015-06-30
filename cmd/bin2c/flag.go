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
		ZF: "zf",
	}
	if s, ok := m[flag]; ok {
		return s
	}
	return fmt.Sprintf("<unknown flag %d>", uint8(flag))
}

const (
	ZF Flag = iota + 1
)

// getFlag converts flag into a corresponding Go expression.
func getFlag(flag Flag) ast.Expr {
	// flagNames maps flag names to their corresponding Go identifiers.
	var flagNames = map[Flag]*ast.Ident{
		ZF: ast.NewIdent("zf"),
	}
	if expr, ok := flagNames[flag]; ok {
		return expr
	}
	log.Fatal(errutil.Newf("unable to lookup identifer for flag %q", flag))
	panic("unreachable")
}
