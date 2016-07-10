## WIP

This project is a *work in progress*. The implementation is *incomplete* and subject to change. The documentation can be inaccurate.

# Experimental

[![GoDoc](https://godoc.org/github.com/decomp/exp?status.svg)](https://godoc.org/github.com/decomp/exp)

This repository contains throw-away prototypes for experimental tools and libraries related to decompilation. They are intended to facilitate understanding, evaluate concepts, validate ideas, and stress-test designs. Once this insight has been gained, they will be removed in favour of high-quality re-implementations.

## restructure2

### Installation

    go get github.com/decomp/exp/cmd/restructure2

## dot2png

### Installation

    go get github.com/decomp/exp/cmd/dot2png

### Example

```bash
$ dot2png *.dot
2015/06/03 16:04:23 Creating: "main_1a.png"
2015/06/03 16:04:23 Creating: "main_1b.png"
2015/06/03 16:04:23 Creating: "main_2a.png"
2015/06/03 16:04:23 Creating: "main_2b.png"
...
2015/06/03 16:04:24 Creating: "main_6a.png"
2015/06/03 16:04:24 Creating: "main_6b.png"
2015/06/03 16:04:24 Creating: "main.png"
```

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
