package form

import (
	"github.com/ANDferraresso/beagleTK/validator"
)

type Field struct {
	Name      string
	Title     string
	MinLength string
	MaxLength string
	Checks    []validator.Check
}

type UI struct {
	Attrs      map[string]string
	Default    string
	Widget     string
	WsUrl      string
	WsCallback string
	Opts       []map[string]string
}

type Form struct {
	Name         string
	Prefix       string
	Required     []string
	DontValidate []string
	FieldsOrder  []string
	Fields       map[string]Field
	UIs          map[string]UI
	Validator    validator.Validator
}
