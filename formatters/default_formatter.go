package formatters

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type DefaultFormatter struct {
	out io.Writer
}

func NewDefaultWriter() *DefaultFormatter {
	return &DefaultFormatter{out: os.Stdout}
}

func (f *DefaultFormatter) FormatData(d interface{}) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if _, err := f.out.Write(b); err != nil {
		return err
	}
	return nil
}

func (f *DefaultFormatter) FormatStringer(d fmt.Stringer) error {
	fmt.Fprintf(f.out, "%s", d.String())
	return nil
}

func (f *DefaultFormatter) FormatArray(d *[]interface{}) error {
	for _, i := range *d {
		if err := f.FormatData(i); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(f.out, "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (f *DefaultFormatter) FormatStrings(d *[]string) error {
	fmt.Fprintf(f.out, "%s\n", strings.Join(*d, "\n"))
	return nil
}
