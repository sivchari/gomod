package gomod

import (
	"context"
	"os"

	"golang.org/x/mod/modfile"
)

type GoModConfig struct {
	GoVersion   string
	ModulePaths []string
}

func Replace(ctx context.Context, gomodConfig *GoModConfig) error {
	for _, modulePath := range gomodConfig.ModulePaths {
		f, err := os.ReadFile(modulePath)
		if err != nil {
			return err
		}
		parsed, err := modfile.Parse("go.mod", f, nil)
		if err != nil {
			return err
		}
		parsed.AddGoStmt(gomodConfig.GoVersion)
		out, err := parsed.Format()
		if err != nil {
			return err
		}
		if err := os.WriteFile(modulePath, out, 0644); err != nil {
			return err
		}
	}
	return nil
}
