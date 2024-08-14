package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/sivchari/commander"
	"github.com/sivchari/gomod/cmd/run"
)

func main() {
	root := commander.NewCommandManager().Build()
	run := commander.NewCommand(&run.Command{})
	root.Register(run)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := root.Run(ctx); err != nil {
		os.Exit(1)
	}
}
