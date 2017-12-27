package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"bldy.build/bldy/cli"
	"bldy.build/build/graph"
	"github.com/pkg/errors"
)

func hash(ctx context.Context) error {

	args := cli.ContextArgs(ctx)
	wd := cli.ContextWD(ctx)
	if len(args) < 1 {
		return fmt.Errorf("build needs at least one argument")
	}
	u, err := url.Parse(args[0])
	if err != nil {
		return errors.Wrap(err, "build")
	}
	g, err := graph.New(wd, u.String())
	if err != nil {
		return errors.Wrap(err, "build")
	}

	fmt.Fprintf(os.Stdout, "hash(%s) = %x\n", g.Root.URL, g.Root.HashNode())
	return nil
}
