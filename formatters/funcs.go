package formatters

import "fmt"

func FormatData(name string, d interface{}) error {
	if f, err := NewFormatter(name); err != nil {
		return err
	} else {
		return f.FormatData(d)
	}
}

func FormatStringer(name string, d fmt.Stringer) error {
	if f, err := NewFormatter(name); err != nil {
		return err
	} else {
		return f.FormatStringer(d)
	}
}

func FormatArray(name string, d *[]interface{}) error {
	if f, err := NewFormatter(name); err != nil {
		return err
	} else {
		return f.FormatArray(d)
	}
}

func FormatStrings(name string, d *[]string) error {
	if f, err := NewFormatter(name); err != nil {
		return err
	} else {
		return f.FormatStrings(d)
	}
}
