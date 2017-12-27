package main

import (
	"context"
	"fmt"
	"net/url"
	"runtime"

	"bldy.build/bldy/cli"
	"bldy.build/bldy/tap"
	"bldy.build/build/builder"
	"bldy.build/build/graph"
	"github.com/pkg/errors"
)

func build(ctx context.Context) error {
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
	b := builder.New(g)
	cpus := int(float32(runtime.NumCPU()) * 1.25)

	ctx, cancel := context.WithCancel(ctx)
	go b.Execute(ctx, cpus)
	display := tap.New()

	go display.Display(b.Updates, cpus)

	for {
		select {
		case done := <-b.Done:
			if done.IsRoot {
				display.Finish()
				cancel()
				return nil
			}

		case err := <-b.Error:
			cancel()
			display.Cancel()
			return err
		case <-b.Timeout:
			cancel()
			return fmt.Errorf("your build has timed out")
		}
	}
}
