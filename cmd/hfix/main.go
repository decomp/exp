// The hfix tool fixes the syntax of IDA generated C header files (*.h -> *.h).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// dbg represents a logger with the "hfix:" prefix, which logs debug messages to
// standard error.
var dbg = log.New(os.Stderr, term.BlueBold("hfix:")+" ", 0)

func usage() {
	const use = `
Fix the syntax of IDA generated C header files (*.h -> *.h).

Usage:

	hfix [OPTION]... FILE.h

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// output specifies the output path.
		output string
		// partial specifies whether to store partially fixed header files.
		partial bool
		// pre specifies whether to store preprocessed header files.
		pre bool
		// quiet specifies whether to suppress non-error messages.
		quiet bool
	)
	flag.StringVar(&output, "o", "", "output path")
	flag.BoolVar(&partial, "partial", false, "store partially fixed header files")
	flag.BoolVar(&pre, "pre", false, "store preprocessed header files")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Parse()
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	hPath := flag.Arg(0)
	// Mute debug messages if `-q` is set.
	if quiet {
		dbg.SetOutput(ioutil.Discard)
	}

	// Read file.
	buf, err := ioutil.ReadFile(hPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Preprocess input.
	input := string(buf)
	input = preprocess(input)
	if pre {
		if err := ioutil.WriteFile("pre.h", []byte(input), 0644); err != nil {
			log.Fatalf("%+v", err)
		}
	}

	// Fix syntax of the IDA generated C header file.
	input, err = fix(input)
	if err != nil {
		if partial {
			if err := ioutil.WriteFile("partial.h", []byte(input), 0644); err != nil {
				log.Fatalf("%+v", err)
			}
		}
		log.Fatalf("%+v", err)
	}

	// Store C header output.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	if _, err := w.WriteString(input); err != nil {
		log.Fatalf("%+v", err)
	}
}

var (
	reEnumSizeSpec = regexp.MustCompile(`(enum [a-zA-Z0-9_$]+) : [a-zA-Z0-9_$]+`)
	reEmptyEnum    = regexp.MustCompile(`enum [a-zA-Z0-9_$]+[\n]{[\n]};[\n]`)
	reAlign        = regexp.MustCompile(`__declspec[(]align[(][0-9]+[)][)] `)
	// Input before:
	//
	//    struct MessageVtbl
	//    {
	//      HRESULT (__stdcall *QueryInterface)(#277 *This, const IID *const riid, void **ppvObject);
	//
	// Input after:
	//
	//    struct MessageVtbl
	//    {
	//      HRESULT (__stdcall *QueryInterface)(MessageVtbl *This, const IID *const riid, void **ppvObject);
	reBrokenTypeRef = regexp.MustCompile(`struct ([a-zA-Z0-9_$]+)[\n]{[\n][^\n#]+(#[0-9]+) [*]This[^\n]+`)
	// Input before:
	//
	//    #pragma pack(push, 8)
	//    #pragma pack(pop)
	//
	// Input after:
	//
	//    empty
	rePragmaPack = regexp.MustCompile(`#pragma pack[(][^)]+[)]`)
	// Input before:
	//
	//	   struct struct_name::$A707B71C060B6D10F73A71917EA8473F::$AA04DEB0C6383F89F13D312A174572A9
	//    {
	//
	// Input after:
	//
	//    empty
	reDupTypeDef = regexp.MustCompile(`[\n](struct|union) ([a-zA-Z0-9_$]+)::[^\n]+[\n]{(.|[\n])+?;[\n][\n]`)
	// Input before:
	//
	//    IDirectDrawClipper::IDirectDrawClipperVtbl
	//
	// Input after:
	//
	//    IDirectDrawClipperVtbl
	reTypeNamespace = regexp.MustCompile(`([a-zA-Z0-9_$]+::)+([a-zA-Z0-9_$]+) `)
	// Input before:
	//
	//    enum enum_name
	//    {
	//      AAA = 0,
	//    };
	//
	// Input after:
	//
	//    enum enum_name
	//    {
	//      AAA = 0,
	//    };
	//
	//    typedef enum enum_name enum_name;
	reTypedefEnum = regexp.MustCompile(`enum ([a-zA-Z0-9_$]+)[\n]{[^}]*};`)
	// Input before:
	//
	//    struct struct_name
	//    {
	//      int x;
	//    };
	//
	// Input after:
	//
	//    struct struct_name
	//    {
	//      int x;
	//    };
	//
	//    typedef struct struct_name struct_name;
	reTypedefStruct = regexp.MustCompile(`struct ([a-zA-Z0-9_$]+)[\n]{(.|[\n])*?[\n]};`)
)

// preprocess fixes simple syntax errors in the given input C header.
func preprocess(input string) string {
	// Drop enum type size specifiers.
	input = reEnumSizeSpec.ReplaceAllString(input, "$1")
	// Remove empty enums.
	input = reEmptyEnum.ReplaceAllString(input, "")
	// Drop alignment attribute.
	input = reAlign.ReplaceAllString(input, "")
	// Drop __unaligned attribute.
	input = strings.Replace(input, "struct __unaligned ", "struct ", -1)
	// Fix broken type names in structs.
	for {
		subs := reBrokenTypeRef.FindAllStringSubmatch(input, 1)
		if subs == nil {
			break
		}
		for _, sub := range subs {
			// struct type name.
			typ := sub[1] + " "
			// #ID
			id := sub[2] + " "
			input = strings.Replace(input, id, typ, -1)
		}
	}
	// Drop #pragma pack directives.
	input = rePragmaPack.ReplaceAllString(input, "")
	// Drop duplicate struct and union type definitions (identified with hash).
	input = reDupTypeDef.ReplaceAllString(input, "\n")
	// Drop namespace in type names.
	input = reTypeNamespace.ReplaceAllString(input, "$2")
	// Insert enum type definitions.
	input = reTypedefEnum.ReplaceAllString(input, "$0\n\ntypedef enum $1 $1;\n")
	// Insert struct type definitions.
	input = reTypedefStruct.ReplaceAllString(input, "$0\n\ntypedef struct $1 $1;\n")
	// Fix syntax of `noreturn` function attributes.
	input = strings.Replace(input, " __noreturn ", " __attribute__((noreturn)) ", -1)
	// Fix destructor method name.
	input = strings.Replace(input, "type_info::`scalar deleting destructor'", "type_info_delete", -1)
	// Fix constructor name.
	input = strings.Replace(input, "type_info::~type_info", "type_info_create", -1)
	return input
}

// fix fixes the syntax of the given IDA generated C header file.
func fix(input string) (string, error) {
	for {
		errbuf := &bytes.Buffer{}
		cmd := exec.Command("clang", "-m32", "-x", "c-header", "-Wno-return-type", "-Wno-invalid-noreturn", "-ferror-limit=0", "-o", "-", "-")
		cmd.Stdin = strings.NewReader(input)
		cmd.Stderr = errbuf
		if err := cmd.Run(); err != nil {
			es, err2 := parseErrors(errbuf.String())
			if err2 != nil {
				return input, errors.WithStack(err2)
			}
			if s, ok := replace(input, es); ok {
				input = s
				// To make it easier to break of an infinite loop, if replacements
				// introduce new Clang errors.
				time.Sleep(1 * time.Millisecond)
				continue
			}
			return input, errors.Wrapf(err, "clang error: %v", errbuf)
		}
		return input, nil
	}
}

// clangError represents an error reported by Clang.
type clangError struct {
	// Line and column number of the error.
	line, col int
	// Error category.
	kind kind
}

// kind represents the set of Clang error categories.
type kind uint

// Clang error categories.
const (
	// error: must use 'struct' tag to refer to type ...
	//
	// Input before:
	//
	//    typedef struct_name type_name;
	//
	// Input after:
	//
	//    typedef struct struct_name type_name;
	kindStructTagMissing kind = iota + 1
	// error: must use 'enum' tag to refer to type
	//
	// Input before:
	//
	//    enum_name foo;
	//
	// Input after:
	//
	//    enum enum_name foo;
	kindEnumTagMissing
	// error: must use 'union' tag to refer to type ...
	//
	// Input before:
	//
	//    typedef union_name type_name;
	//
	// Input after:
	//
	//    typedef union union_name type_name;
	kindUnionTagMissing
	// error: unknown type name '_BYTE'; did you mean 'BYTE'
	//
	// Input before:
	//
	//    _BYTE foo;
	//
	// Input after:
	//
	//    BYTE foo;
	kindByteTypeName
	// error: parameter name omitted
	//
	// Input before:
	//
	//    void f(int, int) {}
	//
	// Input after:
	//
	//    void f(int a1, int a2) {}
	kindParamNameMissing
)

var (
	reError = regexp.MustCompile(`<stdin>:([0-9]+):([0-9]+): (error: [^\n]+)`)
)

// parseErrors parses the error output reported by Clang.
func parseErrors(errbuf string) ([]clangError, error) {
	var es []clangError
	lines := strings.Split(errbuf, "\n")
	for _, line := range lines {
		if !(strings.HasPrefix(line, "<stdin>:") && strings.Contains(line, " error: ")) {
			continue
		}
		subs := reError.FindStringSubmatch(line)
		if subs == nil {
			return nil, errors.Errorf("unable to locate Clang error in line `%v`", line)
		}
		// Parse line number.
		l, err := strconv.Atoi(subs[1])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// Parse column number.
		c, err := strconv.Atoi(subs[2])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		e := clangError{line: l - 1, col: c - 1}
		// Parse error message.
		msg := subs[3]
		switch {
		case strings.HasPrefix(msg, "error: must use 'struct' tag to refer to type"):
			e.kind = kindStructTagMissing
		case strings.HasPrefix(msg, "error: must use 'enum' tag to refer to type"):
			e.kind = kindEnumTagMissing
		case strings.HasPrefix(msg, "error: must use 'union' tag to refer to type"):
			e.kind = kindUnionTagMissing
		case strings.HasPrefix(msg, "error: unknown type name '_BYTE'; did you mean 'BYTE'"):
			e.kind = kindByteTypeName
		case strings.HasPrefix(msg, "error: parameter name omitted"):
			e.kind = kindParamNameMissing
		default:
			// Skip unknown Clang error category.
			continue
			//return nil, errors.Errorf("unable to locate error category for Clang error `%v`", msg)
		}
		es = append(es, e)
	}
	return es, nil
}

// replace fixes the syntax errors identified by Clang in the given input C
// header. The boolean return value indicates that a replacement was made.
func replace(input string, es []clangError) (string, bool) {
	fixed := false
	lines := strings.Split(input, "\n")
	lineFixed := make(map[int]bool)
	for _, e := range es {
		i := e.line
		if lineFixed[i] {
			// Only fix one error per line at the time.
			continue
		}
		line := lines[i]
		switch e.kind {
		case kindStructTagMissing:
			dbg.Printf("replacement made at line %d: kindStructTagMissing", i)
			// insert `struct `
			line = line[:e.col] + "struct " + line[e.col:]
			fixed = true
			lineFixed[i] = true
		case kindEnumTagMissing:
			dbg.Printf("replacement made at line %d: kindEnumTagMissing", i)
			// insert `enum `
			line = line[:e.col] + "enum " + line[e.col:]
			fixed = true
			lineFixed[i] = true
		case kindUnionTagMissing:
			dbg.Printf("replacement made at line %d: kindUnionTagMissing", i)
			// insert `union `
			line = line[:e.col] + "union " + line[e.col:]
			fixed = true
			lineFixed[i] = true
		case kindByteTypeName:
			dbg.Printf("replacement made at line %d: kindByteTypeName", i)
			// replace `_BYTE` with `BYTE`
			line = line[:e.col] + line[e.col+1:]
			fixed = true
			lineFixed[i] = true
		case kindParamNameMissing:
			dbg.Printf("replacement made at line %d: kindParamNameMissing", i)
			// replace `_BYTE` with `BYTE`
			paramName := fmt.Sprintf(" a%d", e.col)
			line = line[:e.col] + paramName + line[e.col:]
			fixed = true
			lineFixed[i] = true
		default:
			panic(fmt.Errorf("support for Clang error kind %v not yet implemented", e.kind))
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n"), fixed
}
