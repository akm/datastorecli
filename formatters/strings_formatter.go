package formatters

import (
	"fmt"
	"io"
	"os"
)

type StringsFormatter struct {
	out io.Writer
}

func NewStringsFormatter() *StringsFormatter {
	return &StringsFormatter{out: os.Stdout}
}

func (f *StringsFormatter) FormatData(d interface{}) error {
	_, err := fmt.Fprintf(f.out, "%v", d)
	return err
}

func (f *StringsFormatter) FormatStringer(d fmt.Stringer) error {
	_, err := fmt.Fprint(f.out, d.String())
	return err
}

func (f *StringsFormatter) FormatArray(d *[]interface{}) error {
	for _, i := range *d {
		if err := f.FormatData(i); err != nil {
			return err
		}
		if _, err := fmt.Fprint(f.out, "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (f *StringsFormatter) FormatStrings(d *[]string) error {
	for _, i := range *d {
		if _, err := fmt.Fprint(f.out, i); err != nil {
			return err
		}
		if _, err := fmt.Fprint(f.out, "\n"); err != nil {
			return err
		}
	}
	return nil
}
