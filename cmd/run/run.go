package run

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/spf13/pflag"

	"github.com/sivchari/gomod"
)

type Command struct{}

func (c *Command) Run(ctx context.Context) error {
	if file == nil || *file == "" {
		panic(errors.New("file is required"))
	}
	f, err := os.ReadFile(*file)
	if err != nil {
		panic(fmt.Errorf("RunCommand read file, but got error: %w", err))
	}
	cfg := &gomod.GoModConfig{}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		panic(fmt.Errorf("RunCommand unmarshal yaml, but got error: %w", err))
	}
	fmt.Println(cfg)
	if err := gomod.Replace(ctx, cfg); err != nil {
		panic(fmt.Errorf("RunCommand tried to replace go version, but got error: %w", err))
	}
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

var file *string

func (c *Command) SetFlags(f *pflag.FlagSet) {
	file = f.StringP("file", "f", "gomod.yaml", "file path")
}
