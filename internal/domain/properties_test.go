package domain

import (
	"bytes"
	"testing"
)

func TestPropertiesWriter_WriteTo(t *testing.T) {
	type fields struct {
		properties map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    string
	}{
		{
			name:   "empty",
			fields: fields{properties: map[string]string{}},
			want:   "",
		},
		{
			name:   "single property",
			fields: fields{properties: map[string]string{"key": "value"}},
			want:   "key=value\n",
		},
		{
			name: "multiple properties",
			fields: fields{properties: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}},
			want: "key1=value1\nkey2=value2\nkey3=value3\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &PropertiesWriter{
				properties: tt.fields.properties,
			}
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
