package domain

import (
	"encoding/json"
	"testing"
)

func TestGeneratorsModelJSON(t *testing.T) {
	tt := []struct {
		name     string
		material string
		generate func(material string) []Asset
		expected []Asset
	}{
		{
			name:     "block model",
			material: "stone",
			generate: func(material string) []Asset {
				bm := BlockModel{All: "minecraft:block/stone"}
				return bm.Generate(material)
			},
			expected: []Asset{
				{
					Name: "stone",
					Data: model{
						Parent:   ParentCubeAll,
						Textures: map[string]string{"all": "minecraft:block/stone"},
					},
				},
			},
		},
		{
			name:     "slab model",
			material: "stone",
			generate: func(material string) []Asset {
				sm := SlabModel{
					Top:    "minecraft:block/stone_top",
					Side:   "minecraft:block/stone_side",
					Bottom: "minecraft:block/stone_bottom",
				}
				return sm.Generate(material)
			},
			expected: []Asset{
				{
					Name: "stone_slab",
					Data: model{
						Parent: ParentSlab,
						Textures: map[string]string{
							"top":    "minecraft:block/stone_top",
							"side":   "minecraft:block/stone_side",
							"bottom": "minecraft:block/stone_bottom",
						},
					},
				},
			},
		},
		{
			name:     "stairs models",
			material: "stone",
			generate: func(material string) []Asset {
				stm := StairsModel{
					Top:    "minecraft:block/stone_top",
					Side:   "minecraft:block/stone_side",
					Bottom: "minecraft:block/stone_bottom",
				}
				return stm.Generate(material)
			},
			expected: []Asset{
				{
					Name: "stone_stairs",
					Data: model{
						Parent: ParentStairs,
						Textures: map[string]string{
							"top":    "minecraft:block/stone_top",
							"side":   "minecraft:block/stone_side",
							"bottom": "minecraft:block/stone_bottom",
						},
					},
				},
				{
					Name: "stone_stairs_inner",
					Data: model{
						Parent: ParentInnerStairs,
						Textures: map[string]string{
							"top":    "minecraft:block/stone_top",
							"side":   "minecraft:block/stone_side",
							"bottom": "minecraft:block/stone_bottom",
						},
					},
				},
				{
					Name: "stone_stairs_outer",
					Data: model{
						Parent: ParentOuterStairs,
						Textures: map[string]string{
							"top":    "minecraft:block/stone_top",
							"side":   "minecraft:block/stone_side",
							"bottom": "minecraft:block/stone_bottom",
						},
					},
				},
			},
		},
		{
			name:     "wall models",
			material: "stone",
			generate: func(material string) []Asset {
				wm := WallModel{Wall: "minecraft:block/stone"}
				return wm.Generate(material)
			},
			expected: []Asset{
				{
					Name: "stone_wall_post",
					Data: model{
						Parent:   ParentWallPost,
						Textures: map[string]string{"wall": "minecraft:block/stone"},
					},
				},
				{
					Name: "stone_wall_side",
					Data: model{
						Parent:   ParentWallSide,
						Textures: map[string]string{"wall": "minecraft:block/stone"},
					},
				},
				{
					Name: "stone_wall_side_tall",
					Data: model{
						Parent:   ParentWallSideTall,
						Textures: map[string]string{"wall": "minecraft:block/stone"},
					},
				},
			},
		},
	}

	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			assets := tt.generate(tt.material)

			if len(assets) != len(tt.expected) {
				t.Fatalf("expected %d assets, got %d", len(tt.expected), len(assets))
			}

			for i, asset := range assets {
				if asset.Name != tt.expected[i].Name {
					t.Errorf("asset %d name mismatch: got %q, want %q", i, asset.Name, tt.expected[i].Name)
				}

				gotJSON, err := json.Marshal(asset.Data)
				if err != nil {
					t.Errorf("failed to marshal asset %d data: %v", i, err)
					continue
				}

				expectedJSON, err := json.Marshal(tt.expected[i].Data)
				if err != nil {
					t.Errorf("failed to marshal expected asset %d data: %v", i, err)
					continue
				}

				if string(gotJSON) != string(expectedJSON) {
					t.Errorf("asset %d JSON mismatch:\n got:  %s\n want: %s", i, string(gotJSON), string(expectedJSON))
				} else {
					t.Logf("asset %d JSON: %s", i, string(gotJSON))
				}
			}
		})
	}
}
