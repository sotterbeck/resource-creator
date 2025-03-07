package service

import (
	"os"
	"path/filepath"
	"resource-creator/internal/domain"
	"testing"
)

func TestAltExporter_Export(t *testing.T) {
	type fields struct {
		material string
		textures []string
	}
	tests := []struct {
		name      string
		fields    fields
		wantFiles []string
		wantErr   bool
	}{
		{
			name: "single texture",
			fields: fields{
				material: "stone",
				textures: []string{"testdata/alternate_sample_1.png"},
			},
			wantFiles: []string{
				"assets/minecraft/blockstates/stone.json",
				"assets/minecraft/models/block/stone/stone_0.json",
				"assets/minecraft/textures/block/stone/stone_0.png",
			},
		},
		{
			name: "multiple textures",
			fields: fields{
				material: "stone",
				textures: []string{"testdata/alternate_sample_1.png", "testdata/alternate_sample_2.png"},
			},
			wantFiles: []string{
				"assets/minecraft/blockstates/stone.json",
				"assets/minecraft/models/block/stone/stone_0.json",
				"assets/minecraft/models/block/stone/stone_1.json",
				"assets/minecraft/textures/block/stone/stone_0.png",
				"assets/minecraft/textures/block/stone/stone_1.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			err := domain.CreateResourcePackFiles(t, tmpDir)
			if err != nil {
				t.Fatal("create resource pack files error:", err)
			}

			ae := &AltExporter{
				material: tt.fields.material,
				textures: tt.fields.textures,
			}

			if err := ae.Export(tmpDir); (err != nil) != tt.wantErr {
				t.Errorf("Export() error = %v, wantErr %v", err, tt.wantErr)
			}

			for _, file := range tt.wantFiles {
				fullPath := filepath.Join(tmpDir, file)
				if _, err := os.Stat(fullPath); os.IsNotExist(err) {
					t.Errorf("Expected file %s does not exist", file)
				}
			}

			logAllFiles(t, tmpDir)
		})
	}
}

func logAllFiles(t *testing.T, root string) {
	t.Helper()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			t.Log(relPath)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Error walking the path %q: %v", root, err)
	}
}
