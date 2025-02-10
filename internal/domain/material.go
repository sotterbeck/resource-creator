package domain

import "strings"

func ParseMaterial(material string) (namespace, name string) {
	if !strings.Contains(material, ":") {
		return material, ""
	}
	parts := strings.Split(material, ":")
	namespace = parts[0]
	name = parts[1]
	return
}
