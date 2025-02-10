package service

import (
	"image"
	"os"
	"path/filepath"
	"resource-creator/internal/domain"
	"testing"
)

func TestCTMExporter_Export(t *testing.T) {
	type fields struct {
		material   string
		tileRes    int
		patternRes domain.Resolution
	}
	tests := []struct {
		name      string
		fields    fields
		wantFiles []string
		wantErr   bool
	}{
		{
			name: "valid",
			fields: fields{
				material:   "minecraft:stone",
				tileRes:    16,
				patternRes: domain.Resolution{Width: 64, Height: 64},
			},
			wantFiles: []string{"stone.properties"},
		},
		{
			name: "invalid tile resolution",
			fields: fields{
				material:   "minecraft:stone",
				tileRes:    15,
				patternRes: domain.Resolution{Width: 64, Height: 64},
			},
			wantErr: true,
		},
		{
			name: "invalid pattern resolution",
			fields: fields{
				material:   "minecraft:stone",
				tileRes:    16,
				patternRes: domain.Resolution{Width: 63, Height: 64},
			},
			wantErr: true,
		},
		{
			name: "tile resolution larger than pattern resolution",
			fields: fields{
				material:   "minecraft:stone",
				tileRes:    64,
				patternRes: domain.Resolution{Width: 16, Height: 16},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			ce := &CTMExporter{
				material:   tt.fields.material,
				tileRes:    tt.fields.tileRes,
				patternRes: tt.fields.patternRes,
			}
			if err := ce.Export(tmpDir); (err != nil) != tt.wantErr {
				t.Errorf("Export() error = %v, wantErr %v", err, tt.wantErr)
			}

			files, err := os.ReadDir(tmpDir)
			if err != nil {
				t.Fatalf("failed to read directory: %v", err)
			}

			if len(files) != len(tt.wantFiles) {
				t.Errorf("Export() got = %d files, want %d files", len(files), len(tt.wantFiles))
			}

			for _, wanted := range tt.wantFiles {
				filePath := filepath.Join(tmpDir, wanted)
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					t.Errorf("Export() file %s not found", wanted)
				}
			}
		})
	}
}

func TestImageExporter_Export(t *testing.T) {
	type fields struct {
		img     image.Image
		tileRes int
	}
	tests := []struct {
		name      string
		fields    fields
		wantFiles []string
	}{
		{
			name: "one split",
			fields: fields{
				img:     image.NewRGBA(image.Rect(0, 0, 8, 8)),
				tileRes: 8,
			},
			wantFiles: []string{"0.png"},
		},
		{
			name: "multiple splits",
			fields: fields{
				img:     image.NewRGBA(image.Rect(0, 0, 8, 8)),
				tileRes: 4,
			},
			wantFiles: []string{"0.png", "1.png", "2.png", "3.png"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			ie := &ImageExporter{
				img:     tt.fields.img,
				tileRes: tt.fields.tileRes,
			}
			if err := ie.Export(tmpDir); err != nil {
				t.Errorf("Export() error = %v", err)
			}

			files, err := os.ReadDir(tmpDir)
			if err != nil {
				t.Fatalf("failed to read directory: %v", err)
			}

			if len(files) != len(tt.wantFiles) {
				t.Errorf("Export() got = %d files, want %d files", len(files), len(tt.wantFiles))
			}

			for _, wanted := range tt.wantFiles {
				filePath := filepath.Join(tmpDir, wanted)
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					t.Errorf("Export() file %s not found", wanted)
				}
			}
		})
	}
}
