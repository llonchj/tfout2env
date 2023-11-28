package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	var filename = flag.String("f", "go.mod", "go mod file path")
	flag.Parse()

	b, err := os.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, modfile.ModulePath(b))
}
