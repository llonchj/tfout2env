package main

import (
	"context"
	"fmt"
	"os"

	"github.com/llonchj/tfout2env"
)

func process(ctx context.Context) error {
	o, err := tfout2env.New(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, o.String())
	return nil
}

func main() {
	ctx := context.Background()
	if err := process(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
