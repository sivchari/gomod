package run

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
)

type Command struct{}

func (c *Command) Run(ctx context.Context) error {
	fmt.Println("Run")
	return nil
}

func (c *Command) Name() string {
	return "run"
}

func (c *Command) Short() string {
	return "Run command"
}

func (c *Command) Long() string {
	return "Run command"
}

func (c *Command) SetFlags(f *pflag.FlagSet) {
	f.String("name", "run", "Name")
}
