package formatters

import "fmt"

type Formatter interface {
	FormatData(d interface{}) error
	FormatStringer(d fmt.Stringer) error
	FormatArray(d *[]interface{}) error
	FormatStrings(d *[]string) error
}
