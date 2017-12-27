package cli

import (
	"context"
	"os"
)

type contextKey int

const (
	ContextKeyArgs contextKey = iota
	ContextKeyBinaryName
	ContextKeyBinaryCommand
	ContextKeyWD
)

// WithArgs adds the remaining args after it's been consumes by Main
func withArgs(parent context.Context, args []string) context.Context {
	return context.WithValue(parent, ContextKeyArgs, args)
}

// ContextArgs returns the cli Args that weren't consumes in main
func ContextArgs(ctx context.Context) []string {
	if args, ok := ctx.Value(ContextKeyArgs).([]string); ok {
		return args
	}
	return nil
}

func withBinaryName(parent context.Context, bin string) context.Context {
	return context.WithValue(parent, ContextKeyBinaryName, bin)
}
func withCommandName(parent context.Context, cmd string) context.Context {
	return context.WithValue(parent, ContextKeyBinaryName, cmd)
}

// WithMeta adds meta information to the context, like working director
func withMeta(parent context.Context) context.Context {
	wd, _ := os.Getwd()
	return context.WithValue(parent, ContextKeyWD, wd)
}

// ContextWD returns the working directory for the context
func ContextWD(ctx context.Context) string {
	if wd, ok := ctx.Value(ContextKeyWD).(string); ok {
		return wd
	}
	return ""
}
