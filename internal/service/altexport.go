package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"resource-creator/internal/domain"
)

type AltExporter struct {
	material string
	textures []string
}

func NewAltExporter(material string, textures []string) Exporter {
	return &AltExporter{
		material: material,
		textures: textures,
	}
}

func (ae *AltExporter) Export(dir string) error {
	var exporters []Exporter
	var models []string

	namespace, mat, err := domain.ParseMaterial(ae.material)
	if err != nil {
		return err
	}

	for i, texturePath := range ae.textures {
		f, err := os.Open(texturePath)
		if err != nil {
			return err
		}

		err = validateTexture(err, f)
		if err != nil {
			return err
		}

		name := fmt.Sprintf("%s_%d", mat, i)

		model := &domain.BlockModel{All: fmt.Sprintf("%s:block/%s/%s", namespace, ae.material, name)}
		modelAsset := model.Generate(name)

		for _, asset := range modelAsset {
			models = append(models, fmt.Sprintf("%s:block/%s/%s", namespace, ae.material, asset.Name))
		}

		exporters = append(exporters, newAssetExporter(modelAsset, domain.AssetTypeModel, namespace, ae.material))
		f.Close()
	}

	blockstate := &domain.CubeBlockState{Models: models}
	blockstateAsset := blockstate.Generate(ae.material)
	exporters = append(exporters, newAssetExporter(blockstateAsset, domain.AssetTypeBlockState, namespace, ae.material))
	exporters = append(exporters, newTextureExporter(ae.textures, namespace, ae.material))

	for _, exporter := range exporters {
		err := exporter.Export(dir)
		if err != nil {
			return fmt.Errorf("failed to export: %w", err)
		}
	}

	return nil
}

func validateTexture(err error, f *os.File) error {
	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("could not decode image: %w", err)
	}

	validator := domain.PatternImage{}
	if !validator.IsValidRes(img.Bounds().Dx(), img.Bounds().Dy()) {
		return errors.New("invalid texture resolution")
	}
	return nil
}

type assetExporter struct {
	typ       string
	namespace string
	material  string
	assets    []domain.Asset
}

func newAssetExporter(assets []domain.Asset, typ, namespace, material string) Exporter {
	return &assetExporter{assets: assets, typ: typ, namespace: namespace, material: material}
}

func (ae *assetExporter) Export(dir string) error {
	path, err := domain.GetResourcePackDir(dir, ae.typ, ae.namespace, ae.material)
	if err != nil {
		return fmt.Errorf("could not get resource pack directory: %w", err)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create resource pack sub directory: %w", err)
	}

	for _, asset := range ae.assets {
		jsonPath := filepath.Join(path, fmt.Sprintf("%s.json", asset.Name))
		file, err := os.Create(jsonPath)
		if err != nil {
			return err
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(asset.Data); err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

type textureExporter struct {
	namespace string
	material  string
	textures  []string
}

func newTextureExporter(textures []string, namespace, material string) Exporter {
	return &textureExporter{textures: textures, namespace: namespace, material: material}
}

func (ae *textureExporter) Export(dir string) error {
	for i, texture := range ae.textures {
		path, err := domain.GetResourcePackDir(dir, domain.AssetTypeTexture, ae.namespace, ae.material)
		if err != nil {
			return err
		}

		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create resource pack sub directory: %w", err)
		}

		source, err := os.Open(texture)
		if err != nil {
			return fmt.Errorf("failed to open texture: %w", err)
		}

		sourceInfo, err := source.Stat()
		if err != nil {
			return fmt.Errorf("failed to retrieve stats of texture: %w", err)
		}

		destinationFile, err := os.OpenFile(filepath.Join(path, fmt.Sprintf("%s_%d.png", ae.material, i)), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, sourceInfo.Mode())
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}

		_, err = io.Copy(destinationFile, source)
		if err != nil {
			return fmt.Errorf("failed to copy contents: %w", err)
		}

		err = destinationFile.Sync()
		if err != nil {
			return fmt.Errorf("failed to sync destination file: %w", err)
		}

		source.Close()
		destinationFile.Close()
	}
	return nil
}
