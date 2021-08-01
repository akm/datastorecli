package formatters

import (
	"fmt"
)

type NdjsonFormatter struct {
	JsonFormatter
}

func NewNdjsonFormatter() *NdjsonFormatter {
	return &NdjsonFormatter{JsonFormatter: *NewJsonFormatter()}
}

func (f *NdjsonFormatter) FormatArray(d *[]interface{}) error {
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

func (f *NdjsonFormatter) FormatStrings(d *[]string) error {
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
