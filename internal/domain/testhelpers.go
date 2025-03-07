package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func CreateResourcePackFiles(t *testing.T, tmpDir string) error {
	t.Helper()

	f, err := os.Create(filepath.Join(tmpDir, "pack.mcmeta"))
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(tmpDir, "assets"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()
	return err
}
