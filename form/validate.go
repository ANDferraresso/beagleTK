package form

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/valyala/fasthttp"
)

//
func (form *Form) ValidateField(k string, value string) bool {
	fmt.Println("Validate:", k, value)
	if stringInSlice(k, form.DontValidate) {
		return true
	}

	if form.Fields[k].MinLength != "" {
		length, err := strconv.Atoi(form.Fields[k].MinLength)
		if err != nil {
			return false
		}
		if utf8.RuneCountInString(value) < length {
			return false
		}
	}

	if form.Fields[k].MaxLength != "" {
		length, err := strconv.Atoi(form.Fields[k].MaxLength)
		if err != nil {
			return false
		}
		if utf8.RuneCountInString(value) > length {
			return false
		}
	}

	// Checks
	for _, check := range form.Fields[k].Checks {
		if check.Func != "" {
			if !(form.Validator.Validate(check.Func, value, &check.Pars)) {
				return false
			}
		}
	}

	return true
}

//
func (form *Form) ValidateAllCtx(ctx *fasthttp.RequestCtx) (bool, map[string]string, []string) {
	// https://godoc.org/github.com/valyala/fasthttp#RequestCtx.FormValue
	// https://stackoverflow.com/questions/39265978/get-a-request-parameter-key-value-in-fasthttp

	fValues := map[string]string{}
	wrongFields := []string{}

	fmt.Println("form.FieldsOrder:", form.FieldsOrder)
	for _, k := range form.FieldsOrder {
		if k == "_csrf" {

		} else {
			fmt.Println("Sono qui (A):", k)
			if !ctx.PostArgs().Has(form.Prefix + k) {
				// Parametro non presente.
				fmt.Println("Parametro non presente:", form.Prefix + k)
				fValues[k] = ""
				if !stringInSlice(k, form.DontValidate) {
					wrongFields = append(wrongFields, k)
				}
			} else {
				fValues[k] = strings.Trim(string(ctx.PostArgs().Peek(form.Prefix+k)), " ")
			    fmt.Println("fValues[k]:", fValues[k])
				// Se la lunghezza minima consentita è "0" (o "") e l'input ha lunghezza 0, non lo valida.
				if (form.Fields[k].MinLength == "" || form.Fields[k].MinLength == "0") && utf8.RuneCountInString(fValues[k]) == 0 {
					// Se la lunghezza minima consentita è 0 (o null) e l'input ha lunghezza 0, non lo validare.
				} else {
					if !form.ValidateField(k, fValues[k]) {
						wrongFields = append(wrongFields, k)
					}
				}
			}
		}
	}

	if len(wrongFields) > 0 {
		return false, fValues, wrongFields
	}

	return true, fValues, wrongFields
}

//
func (form *Form) ValidateAllMap(POST_ map[string]string) (bool, map[string]string, []string) {
	fValues := map[string]string{}
	wrongFields := []string{}

	fmt.Println("form.FieldsOrder:", form.FieldsOrder)
	for _, k := range form.FieldsOrder {
		if k == "_csrf" {
			;
		} else {
            val, ok := POST_[form.Prefix + k]
			if !ok {
			    // Parametro non presente.
				fmt.Println("Parametro non presente:", form.Prefix + k)
				fValues[k] = ""
				if !stringInSlice(k, form.DontValidate) {
					wrongFields = append(wrongFields, k)
				}		
			} else {
				fValues[k] = strings.Trim(val, " ")
			    fmt.Println("fValues[k]:", fValues[k])
				// Se la lunghezza minima consentita è "0" (o "") e l'input ha lunghezza 0, non lo valida.
				if (form.Fields[k].MinLength == "" || form.Fields[k].MinLength == "0") && utf8.RuneCountInString(fValues[k]) == 0 {
					// Se la lunghezza minima consentita è 0 (o null) e l'input ha lunghezza 0, non lo validare.
				} else {
					if !form.ValidateField(k, fValues[k]) {
						wrongFields = append(wrongFields, k)
					}
				}
			}
		}
	}

	if len(wrongFields) > 0 {
		return false, fValues, wrongFields
	}

	return true, fValues, wrongFields
}