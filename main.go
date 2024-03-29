package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/sonatard/proto-gen-swiftapiclient/printer"
	"golang.org/x/xerrors"
)

func main() {
	opt, paths, err := parseOption()
	if err != nil || len(paths) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(paths, opt.importPaths); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
}

func run(files []string, importPaths []string) error {
	p := protoparse.Parser{
		ImportPaths: importPaths,
	}

	fds, err := p.ParseFiles(files...)
	if err != nil {
		return xerrors.Errorf("Unable to parse pb file: %v \n", err)
	}

	printer.Print(fds)

	return nil
}
