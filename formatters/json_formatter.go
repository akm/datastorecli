package formatters

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JsonFormatter struct {
	out         io.Writer
	marshalFunc func(v interface{}) ([]byte, error)
}

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{out: os.Stdout, marshalFunc: json.Marshal}
}

func NewPrettyJsonFormatter() *JsonFormatter {
	return &JsonFormatter{
		out: os.Stdout,
		marshalFunc: func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "  ")
		},
	}
}

func (f *JsonFormatter) FormatData(d interface{}) error {
	b, err := f.marshalFunc(d)
	if err != nil {
		return err
	}
	if _, err := f.out.Write(b); err != nil {
		return err
	}
	return nil
}

func (f *JsonFormatter) FormatStringer(d fmt.Stringer) error {
	return f.FormatData(d.String())
}

func (f *JsonFormatter) FormatArray(d *[]interface{}) error {
	return f.FormatData(*d)
}

func (f *JsonFormatter) FormatStrings(d *[]string) error {
	return f.FormatData(*d)
}
