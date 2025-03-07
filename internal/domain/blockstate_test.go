package domain

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type blockStateGenerator interface {
	Generate(material string) []Asset
}

func loadFixture(t *testing.T, filename string) string {
	t.Helper()
	path := filepath.Join("testdata", "blockstate", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read fixture file %s: %v", path, err)
	}
	return string(data)
}

func TestBlockStateGenerate(t *testing.T) {
	tests := []struct {
		name      string
		generator blockStateGenerator
		material  string
		wantName  string
		wantFile  string
	}{
		{
			name:      "cube block state",
			generator: &CubeBlockState{Models: []string{"minecraft:block/stone1", "minecraft:block/stone2"}},
			material:  "stone",
			wantName:  "stone",
			wantFile:  "cube_fixture.json",
		},
		{
			name:      "slab block state",
			generator: &SlabBlockState{BaseModels: []string{"minecraft:block/stone1", "minecraft:block/stone2"}},
			material:  "stone",
			wantName:  "stone_slab",
			wantFile:  "slab_fixture.json",
		},
		{
			name:      "stairs block state single variant",
			generator: &StairsBlockState{BaseModels: []string{"minecraft:block/stone"}},
			material:  "stone",
			wantName:  "stone_stairs",
			wantFile:  "stairs_single_fixture.json",
		},
		{
			name:      "stairs block state multiple variants",
			generator: &StairsBlockState{BaseModels: []string{"minecraft:block/stone1", "minecraft:block/stone2"}},
			material:  "stone",
			wantName:  "stone_stairs",
			wantFile:  "stairs_multiple_fixture.json",
		},
		{
			name:      "wall block state",
			generator: &WallBlockState{BaseModels: []string{"minecraft:block/stone1", "minecraft:block/stone2"}},
			material:  "stone",
			wantName:  "stone_wall",
			wantFile:  "wall_fixture.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assets := tt.generator.Generate(tt.material)
			actual, err := json.MarshalIndent(assets[0].Data, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal assets: %v", err)
			}

			if assets[0].Name != tt.wantName {
				t.Errorf("unexpected asset name\nGot: %s\nWant: %s", assets[0].Name, tt.wantName)
			}

			expected := loadFixture(t, tt.wantFile)

			if string(actual) != expected {
				t.Errorf("unexpected JSON output\nGot:\n%s\nWant:\n%s", string(actual), expected)
			}
		})
	}
}
