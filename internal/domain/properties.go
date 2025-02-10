package domain

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

type PropertiesWriter struct {
	properties map[string]string
}

func NewPropertiesWriter(props map[string]string) *PropertiesWriter {
	return &PropertiesWriter{properties: props}
}

func (pw *PropertiesWriter) WriteTo(w io.Writer) (n int64, err error) {
	var buf bytes.Buffer

	keys := make([]string, 0, len(pw.properties))
	for k := range pw.properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		_, err := fmt.Fprintf(&buf, "%s=%s\n", k, pw.properties[k])
		if err != nil {
			return 0, err
		}
	}
	n, err = buf.WriteTo(w)
	return n, err
}
