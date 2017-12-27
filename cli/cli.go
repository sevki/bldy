package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var l = log.New(os.Stdout, "", 0)

// App represents a cli application
type App struct {
	Name     string
	Version  string
	Authors  []string
	Commands []Command
}

// Command represents a command that the cli app might recieve
type Command struct {
	Name    string
	Command string
	Aliases []string
	Usage   string
	Action  Action
}

// Action represents a cli commands funcion when it's invoked
type Action func(context.Context) error

// Main is the entry point for your app
func (a *App) Main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	exit := func(err error) {
		cancel()
		if err != nil {
			l.Println(err)
			l.Println()
			helpStd(a)
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	}
	args := os.Args // list args
	bin := args[0]  // binary name
	_ = bin         // suppress the unused error
	args = args[1:] // get rid of the binary name
	if len(args) < 1 {
		e := errors.New("we need at least 1 argument, you have none")
		exit(e)
	}
	cmdName := args[0] // second arg is a command

	if strings.HasPrefix(cmdName, "//") {
		cmdName = "build"
	} else {
		args = args[1:] // get rid of the command flag
	}

	var c *Command
	for i := range a.Commands {
		if a.Commands[i].Name == cmdName {
			c = &a.Commands[i]
		}
	}

	if c == nil {
		e := fmt.Errorf("%s is not a valid command", cmdName)
		exit(e)
	}

	ctx = withArgs(ctx, args)
	ctx = withBinaryName(ctx, bin)
	ctx = withCommandName(ctx, cmdName)
	ctx = withMeta(ctx)

	err := c.Action(ctx)

	if err != nil {
		exit(err)
	}
	exit(nil)
}
