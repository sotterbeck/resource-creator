package domain

// JSONGenerator is an interface for generating JSON data for resource pack assets.
// Assets can be block states or models.
type JSONGenerator interface {

	// Generate generates JSON data for a specific asset.
	// Returns either a single JSON object or a slice of JSON objects depending on the asset.
	Generate(material string) []Asset
}

type Asset struct {
	Data interface{}
	Name string
}
