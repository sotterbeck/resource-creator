package domain

import (
	"bytes"
	"testing"
)

func TestPropertiesWriter_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		writer  func() *PropertiesWriter
		wantErr bool
		want    string
	}{
		{
			name:   "empty",
			writer: func() *PropertiesWriter { return NewPropertiesWriter(map[string]string{}) },
			want:   "",
		},
		{
			name:   "single property",
			writer: func() *PropertiesWriter { return NewPropertiesWriter(map[string]string{"key": "value"}) },
			want:   "key=value\n",
		},
		{
			name: "multiple properties",
			writer: func() *PropertiesWriter {
				return NewPropertiesWriter(map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				})
			},
			want: "key1=value1\nkey2=value2\nkey3=value3\n",
		},
		{
			name: "with comments",
			writer: func() *PropertiesWriter {
				pw := NewPropertiesWriter(map[string]string{"key": "value"})
				pw.AddComment("comment")
				return pw
			},
			want: "# comment\nkey=value\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := tt.writer()
			w := &bytes.Buffer{}
			_, err := pw.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if w.String() != tt.want {
				t.Errorf("WriteTo() got = %v, want %v", w.String(), tt.want)
			}
		})
	}
}
