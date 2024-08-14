package gomod

import (
	"context"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type GoModConfig struct {
	GoVersion   string   `yaml:"go_version"`
	ModulePaths []string `yaml:"module_paths"`
}

func Replace(ctx context.Context, gomodConfig *GoModConfig) error {
	for _, modulePath := range gomodConfig.ModulePaths {
		path := filepath.Join(modulePath, "go.mod")
		f, err := os.ReadFile(path)
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
		if err := os.WriteFile(path, out, 0644); err != nil {
			return err
		}
	}
	return nil
}
