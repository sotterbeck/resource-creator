package domain

import (
	"fmt"
	"strings"
)

type variant struct {
	Model  string `json:"model"`
	UVLock bool   `json:"uvlock,omitempty"`
	X      int    `json:"x,omitempty"`
	Y      int    `json:"y,omitempty"`
}

type multipartCase struct {
	Apply []variant         `json:"apply,omitempty"`
	When  map[string]string `json:"when,omitempty"`
}

type blockState struct {
	Variants  map[string][]variant `json:"variants,omitempty"`
	Multipart []multipartCase      `json:"multipart,omitempty"`
}

type modelType struct {
	typ    string
	suffix string
}

func newVariant(models []string) []variant {
	return newRotatedVariant(models, Rotation{})
}

func newRotatedVariant(models []string, rot Rotation) []variant {
	var variants []variant
	for _, m := range models {
		variants = append(variants, variant{Model: m, UVLock: rot.UVLock, X: rot.X, Y: rot.Y})
	}
	return variants
}

func withSuffix(models []string, suffix string) []string {
	result := make([]string, len(models))
	for i, m := range models {
		result[i] = m + suffix
	}
	return result
}

type CubeBlockState struct {
	Models []string
}

func (bs *CubeBlockState) Generate(material string) []Asset {
	return []Asset{
		{
			Name: material,
			Data: blockState{
				Variants: map[string][]variant{
					"": newVariant(bs.Models),
				},
			},
		},
	}
}

type SlabBlockState struct {
	BaseModels []string
}

func (bs *SlabBlockState) Generate(material string) []Asset {
	types := []modelType{
		{"top", "_slab_top"},
		{"bottom", "_slab"},
		{"double", ""},
	}
	variants := make(map[string][]variant)
	for _, t := range types {
		variants[fmt.Sprintf("type=%s", t.typ)] = newVariant(withSuffix(bs.BaseModels, t.suffix))
	}
	return []Asset{
		{
			Name: fmt.Sprintf("%s_slab", material),
			Data: blockState{
				Variants: variants,
			},
		},
	}
}

type StairsBlockState struct {
	BaseModels []string
}

func buildStairModel(baseModel, shape string) string {
	if strings.Contains(shape, "inner") {
		return baseModel + "_stairs_inner"
	}
	if strings.Contains(shape, "outer") {
		return baseModel + "_stairs_outer"
	}
	return baseModel + "_stairs"
}

func (bs *StairsBlockState) Generate(material string) []Asset {
	facing := []string{"north", "east", "south", "west"}
	half := []string{"bottom", "top"}
	shape := []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}

	variants := make(map[string][]variant)
	for _, f := range facing {
		for _, h := range half {
			for _, s := range shape {
				key := fmt.Sprintf("facing=%s,half=%s,shape=%s", f, h, s)
				rot := GetStairRotation(f, h, s)

				models := make([]string, len(bs.BaseModels))
				for i, m := range bs.BaseModels {
					models[i] = buildStairModel(m, s)
				}

				variants[key] = newRotatedVariant(models, rot)
			}
		}
	}
	return []Asset{
		{
			Name: fmt.Sprintf("%s_stairs", material),
			Data: blockState{
				Variants: variants,
			},
		},
	}
}

type WallBlockState struct {
	BaseModels []string
}

func (bs *WallBlockState) Generate(material string) []Asset {
	var cases []multipartCase
	postModels := withSuffix(bs.BaseModels, "_wall_post")
	cases = append(cases, multipartCase{
		Apply: newVariant(postModels),
		When:  map[string]string{"up": "true"},
	})

	directions := []string{"north", "east", "south", "west"}
	types := []modelType{
		{"low", "_wall_side"},
		{"tall", "_wall_side_tall"},
	}

	for _, wallType := range types {
		for _, dir := range directions {
			rot := GetWallRotation(dir)
			sideModels := withSuffix(bs.BaseModels, wallType.suffix)
			cases = append(cases, multipartCase{
				Apply: newRotatedVariant(sideModels, rot),
				When:  map[string]string{dir: wallType.typ},
			})
		}
	}

	return []Asset{
		{
			Name: fmt.Sprintf("%s_wall", material),
			Data: blockState{
				Multipart: cases,
			},
		},
	}
}
