package domain

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

type PropertiesWriter struct {
	properties map[string]string
	comments   []string
}

func NewPropertiesWriter(props map[string]string) *PropertiesWriter {
	return &PropertiesWriter{properties: props}
}

func (pw *PropertiesWriter) AddComment(comment string) {
	pw.comments = append(pw.comments, comment)
}

func (pw *PropertiesWriter) WriteTo(w io.Writer) (n int64, err error) {
	var buf bytes.Buffer

	for _, c := range pw.comments {
		_, err := fmt.Fprintf(&buf, "# %s\n", c)
		if err != nil {
			return 0, err
		}
	}

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
