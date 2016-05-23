## WIP

This project is a *work in progress*. The implementation is *incomplete* and subject to change. The documentation can be inaccurate.

# ll2dot

[![GoDoc](https://godoc.org/github.com/decomp/exp/cmd/ll2dot?status.svg)](https://godoc.org/github.com/decomp/exp/cmd/ll2dot)

`ll2dot` is a tool which generates control flow graphs from LLVM IR assembly files (e.g. *.ll -> *.dot). The output is a set of Graphviz DOT files, each representing the control flow graph of a function using one node per basic block.

For a source file "foo.ll" containing the functions "bar" and "baz" the following DOT files will be generated:

   * foo_graphs/bar.dot
   * foo_graphs/baz.dot

## Installation

```shell
go get github.com/decomp/exp/cmd/ll2dot
```

## Usage

```shell
ll2dot [OPTION]... FILE...

Flags:
  -f    force overwrite existing graph directories
  -funcs string
        comma separated list of functions to parse (e.g. "foo,bar")
  -img
        generate an image representation of the CFG
  -q    suppress non-error messages
```

## Examples

### funcs

```shell
ll2dot -f -img -funcs="foo,bar" testdata/funcs.ll
```

Input:
* [funcs.ll](testdata/funcs.ll)

Output:
* [foo.dot](testdata/funcs_graphs/foo.dot)
* [bar.dot](testdata/funcs_graphs/bar.dot)

![CFG funcs the foo function of funcs.ll](https://raw.githubusercontent.com/decomp/ll2dot/master/testdata/funcs_graphs/foo.png)
![CFG funcs the bar function of funcs.ll](https://raw.githubusercontent.com/decomp/ll2dot/master/testdata/funcs_graphs/bar.png)

### switch

```shell
ll2dot -f -img testdata/funcs.ll
```

Input:
* [switch.ll](testdata/switch.ll)

Output:
* [main.dot](testdata/switch_graphs/main.dot)

![CFG switch the main function of switch.ll](https://raw.githubusercontent.com/decomp/ll2dot/master/testdata/switch_graphs/main.png)

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
