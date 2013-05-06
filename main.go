package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"sort"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s: %s <go-src>...\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	fset := token.NewFileSet()
	imports := make(map[string]struct{})

	for _, fname := range flag.Args() {
		f, err := parser.ParseFile(fset, fname, nil, parser.ImportsOnly)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't parse %s: %v\n", fname, err)
			os.Exit(1)
		}

		for _, imp := range f.Imports {
			imports[imp.Path.Value] = struct{}{}
		}
	}

	importList := make([]string, 0, len(imports))

	for imp, _ := range imports {
		importList = append(importList, imp[1:len(imp)-1])
	}

	sort.StringSlice(importList).Sort()

	for _, imp := range importList {
		fmt.Println(imp)
	}
}
