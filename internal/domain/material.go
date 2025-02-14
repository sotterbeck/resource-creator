package domain

import (
	"errors"
	"strings"
)

func ParseMaterial(material string) (namespace, name string, err error) {
	if strings.Count(material, ":") > 1 {
		return "", "", errors.New("material contains more than one colon")
	}

	if material == "" {
		return "", "", nil
	}

	if !strings.Contains(material, ":") {
		return "minecraft", material, nil
	}
	parts := strings.Split(material, ":")
	namespace = parts[0]
	name = parts[1]
	return
}
