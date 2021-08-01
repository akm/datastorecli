package formatters

import "github.com/pkg/errors"

func NewFormatter(name string) (Formatter, error) {
	switch name {
	case "json":
		return NewJsonFormatter(), nil
	case "pretty-json":
		return NewPrettyJsonFormatter(), nil
	case "ndjson":
		return NewNdjsonFormatter(), nil
	case "strings":
		return NewStringsFormatter(), nil
	default:
		return nil, errors.Errorf("Unsupported formatter %q", name)
	}
}
