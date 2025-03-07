package domain

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	AssetTypeModel      = "model"
	AssetTypeBlockState = "blockstate"
	AssetTypeTexture    = "texture"
)

// IsResourcePack checks whether the given directory is a valid resource pack directory.
func IsResourcePack(dir string) (bool, error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	if !dirInfo.IsDir() {
		return false, fmt.Errorf("%s is not a directory", dir)
	}

	packMetaPath := filepath.Join(dir, "pack.mcmeta")
	packMetaInfo, err := os.Stat(packMetaPath)
	if err != nil || packMetaInfo.IsDir() {
		return false, nil
	}

	assetPath := filepath.Join(dir, "assets")
	assetInfo, err := os.Stat(assetPath)
	if err != nil || !assetInfo.IsDir() {
		return false, nil
	}

	return true, nil
}

// GetResourcePackDir returns the directory path for a specific resource type within a resource pack.
//
// The function first validates if the provided baseDir corresponds to a resource pack using IsResourcePack.
// If the baseDir is not a valid resource pack directory, it returns an error.
func GetResourcePackDir(baseDir, typ, namespace, material string) (string, error) {
	isPack, err := IsResourcePack(baseDir)
	if err != nil {
		return "", err
	}

	if !isPack {
		return "", fmt.Errorf("base dir %s is not a resourcepack directory", baseDir)
	}

	switch typ {
	case AssetTypeModel:
		return filepath.Join(baseDir, "assets", namespace, "models", "block", material), nil
	case AssetTypeBlockState:
		return filepath.Join(baseDir, "assets", namespace, "blockstates"), nil
	case AssetTypeTexture:
		return filepath.Join(baseDir, "assets", namespace, "textures", "block", material), nil
	}

	return "", fmt.Errorf("unknown resource pack type: %s", typ)
}
