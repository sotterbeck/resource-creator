package domain

import (
	"fmt"
)

const (
	ParentCubeAll      = "minecraft:block/cube_all"
	ParentSlab         = "minecraft:block/slab"
	ParentStairs       = "minecraft:block/stairs"
	ParentInnerStairs  = "minecraft:block/inner_stairs"
	ParentOuterStairs  = "minecraft:block/outer_stairs"
	ParentWallPost     = "minecraft:block/template_wall_post"
	ParentWallSide     = "minecraft:block/template_wall_side"
	ParentWallSideTall = "minecraft:block/template_wall_side_tall"
)

type model struct {
	Parent   string            `json:"parent"`
	Textures map[string]string `json:"textures"`
}

type BlockModel struct {
	All string
}

func (m *BlockModel) Generate(material string) []Asset {
	return []Asset{{
		Name: material,
		Data: model{ParentCubeAll, map[string]string{
			"all": m.All,
		}},
	}}
}

type SlabModel struct {
	Top    string
	Side   string
	Bottom string
}

func (m *SlabModel) Generate(material string) []Asset {
	return []Asset{{
		Name: fmt.Sprintf("%s_slab", material),
		Data: model{ParentSlab, map[string]string{
			"top":    m.Top,
			"side":   m.Side,
			"bottom": m.Bottom,
		}},
	},
	}
}

type StairsModel struct {
	Top    string
	Side   string
	Bottom string
}

func (m *StairsModel) Generate(material string) []Asset {
	models := make([]Asset, 0)
	parents := []string{ParentStairs, ParentInnerStairs, ParentOuterStairs}
	for _, parent := range parents {
		suffix := ""
		switch parent {
		case ParentInnerStairs:
			suffix = "stairs_inner"
		case ParentOuterStairs:
			suffix = "stairs_outer"
		case ParentStairs:
			suffix = "stairs"
		}

		models = append(models, Asset{
			Name: fmt.Sprintf("%s_%s", material, suffix),
			Data: model{parent, map[string]string{
				"top":    m.Top,
				"side":   m.Side,
				"bottom": m.Bottom,
			}},
		})
	}
	return models
}

type WallModel struct {
	Wall string
}

func (m *WallModel) Generate(material string) []Asset {
	models := make([]Asset, 0)
	parents := []string{ParentWallPost, ParentWallSide, ParentWallSideTall}

	for _, parent := range parents {
		suffix := ""
		switch parent {
		case ParentWallPost:
			suffix = "wall_post"
		case ParentWallSide:
			suffix = "wall_side"
		case ParentWallSideTall:
			suffix = "wall_side_tall"
		}

		models = append(models, Asset{
			Name: fmt.Sprintf("%s_%s", material, suffix),
			Data: model{parent, map[string]string{
				"wall": m.Wall,
			}},
		})
	}
	return models
}
