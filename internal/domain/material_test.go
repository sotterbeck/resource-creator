package domain

import "testing"

func TestParseMaterial(t *testing.T) {
	tests := []struct {
		name          string
		material      string
		wantNamespace string
		wantName      string
		wantErr       bool
	}{
		{
			name:          "material with explicit namespace",
			material:      "namespace:name",
			wantNamespace: "namespace",
			wantName:      "name",
		},
		{
			name:          "material with implicit namespace",
			material:      "name",
			wantNamespace: "minecraft",
			wantName:      "name",
		},
		{
			name:          "empty material",
			material:      "",
			wantNamespace: "",
			wantName:      "",
		},
		{
			name:     "material with more than one colon",
			material: "namespace:name:extra",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamespace, gotName, err := ParseMaterial(tt.material)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNamespace != tt.wantNamespace {
				t.Errorf("ParseMaterial() gotNamespace = %v, want %v", gotNamespace, tt.wantNamespace)
			}
			if gotName != tt.wantName {
				t.Errorf("ParseMaterial() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
