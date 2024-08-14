package gomod

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReplace(t *testing.T) {
	want := `module testdata

go 1.17
`
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	path := filepath.Join(wd, "testdata/go.mod")

	if err := Replace(context.Background(), &GoModConfig{
		ModulePaths: []string{path},
		GoVersion:   "1.17",
	}); err != nil {
		t.Fatalf("error: %v", err)
	}

	t.Cleanup(func() {
		Replace(context.Background(), &GoModConfig{
			ModulePaths: []string{path},
			GoVersion:   "1.16",
		})
	})

	f, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if diff := cmp.Diff(string(f), want); diff != "" {
		t.Fatalf("(-got, +want)\n%s", diff)
	}
}
