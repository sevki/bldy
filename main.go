// Copyright 2015-2016 Sevki <s@sevki.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bldy.build/bldy/cli"

	"flag"
)

var (
	buildVer = "version"
)
var (
	disp    = flag.String("d", "tap", "only available display is tap currently")
	display Display
)

func main() {
	app := cli.App{}
	app.Name = "bldy"
	app.Version = buildVer
	app.Commands = []cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build a target",
			Action:  build,
		},
		{
			Name:    "hash",
			Aliases: []string{"h"},
			Usage:   "hash a target",
			Action:  hash,
		},
	}
	app.Main()
}
