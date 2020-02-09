// Copyright 2018 Sevki <s@sevki.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"flag"
	"os"

<<<<<<< HEAD
	"bldy.build/bldy/cmd/build"
	"bldy.build/bldy/cmd/trace"
	"bldy.build/bldy/fileutils"
=======
	"bldy.build/bldy/src/cmd/build"
<<<<<<< HEAD:src2/cmd/bldy/main.go
=======
	"bldy.build/bldy/src/cmd/query"
>>>>>>> 97e98155e24d9c7de236ebaf33a5557c36660e2d:src/cmd/bldy/main.go
	"bldy.build/bldy/src/url"
>>>>>>> 97e98155e24d9c7de236ebaf33a5557c36660e2d
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&build.BuildCmd{}, "")
<<<<<<< HEAD
	subcommands.Register(&trace.TraceCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	if u, err := fileutils.ResolveFromWD(flag.Arg(1)); err == nil {
=======

	flag.Parse()
	ctx := context.Background()
	if u, err := url.Parse(flag.Arg(1)); err == nil {
>>>>>>> 97e98155e24d9c7de236ebaf33a5557c36660e2d
		os.Exit(int(subcommands.Execute(ctx, u)))
	} else {
		os.Exit(int(subcommands.Execute(ctx)))
	}
}
