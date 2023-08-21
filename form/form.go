package form

import (

	//	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/valyala/fasthttp"

	"github.com/ANDferraresso/beagleTK/orm"
	"github.com/ANDferraresso/beagleTK/validator"
)

func (form *Form) SetupForm() {
	form.Name = ""
	form.Prefix = ""
	form.Required = []string{}
	form.DontValidate = []string{}
	form.FieldsOrder = []string{}
	form.Fields = make(map[string]*Field)
	form.UIs = make(map[string]*UI)
	form.Validator = validator.Validator{}
	form.Validator.SetupValidator()
}

func (form *Form) AddField(fd Field) {
	form.FieldsOrder = append(form.FieldsOrder, fd.Name)
	form.Fields[fd.Name] = &fd
	form.UIs[fd.Name] = &UI{
		Attrs:      make(map[string]string),
		Default:    "",
		Widget:     "",
		WsUrl:      "",
		WsCallback: "",
	}
}

func (form *Form) AddCsrfField(value string) {
	form.FieldsOrder = append(form.FieldsOrder, "_csrf")
	form.Fields["_csrf"] = &Field{
		Name:      "_csrf",
		Title:     "",
		MinLength: "",
		MaxLength: "",
		Checks:    []validator.Check{},
	}
	form.UIs["_csrf"] = &UI{
		Attrs:      make(map[string]string),
		Default:    "",
		Widget:     "",
		WsUrl:      "",
		WsCallback: "",
	}

	form.UIs["_csrf"] = &UI{Attrs: map[string]string{"value": value}}
}

/*
    // Setta la sessione anti CSRF
    public function setCsrfCookie($value) {
        list($mSec, $sec) = explode(" ", microtime());
        if (isset($_SESSION[CSRF_PREFIX.'_csrf'])) {
            // Esiste già
            if (!is_array($_SESSION[CSRF_PREFIX.'_csrf']) || empty($_SESSION[CSRF_PREFIX.'_csrf'])) { // Verifica se è vuoto oppure NULL
                $_SESSION[CSRF_PREFIX.'_csrf'] = array();
            }

            // Per evitare eventuali improbabili collisioni
            if (!array_key_exists($value, $_SESSION[CSRF_PREFIX.'_csrf'])) {
                // Non esiste
                // Se ci sono troppi token (cosa sospetta...)
                if (count($_SESSION[CSRF_PREFIX.'_csrf']) >= 10) {
                    $_SESSION[CSRF_PREFIX.'_csrf'] = array_slice($_SESSION[CSRF_PREFIX.'_csrf'], 1);
                    $_SESSION[CSRF_PREFIX.'_csrf'][$value] = (string)((float)$sec + (float)$mSec + CRSF_TOKEN_DURATION);
                }
                else {
                    $_SESSION[CSRF_PREFIX.'_csrf'][$value] = (string)((float)$sec + (float)$mSec + CRSF_TOKEN_DURATION);
                }
            }
        }
        else {
            // Non esiste
            $_SESSION[CSRF_PREFIX.'_csrf'] = array();
            $_SESSION[CSRF_PREFIX.'_csrf'][$value] = (string)((float)$sec + (float)$mSec + CRSF_TOKEN_DURATION);
        }
    }

	func (form *Form) ValidateCsrfToken(value string) {

	}

    //
    public function validateCsrfToken() {
        if ((!isset($_POST[$this->fDef['prefix'].'_csrf']) && !isset($_GET[$this->fDef['prefix'].'_csrf']))
            || !isset($_SESSION[CSRF_PREFIX.'_csrf']) || !is_array($_SESSION[CSRF_PREFIX.'_csrf'])) {
            return FALSE;
        }
        else {
            $token = "";
            if (isset($_POST[$this->fDef['prefix'].'_csrf']))
                $token = $_POST[$this->fDef['prefix'].'_csrf'];
            else if (isset($_GET[$this->fDef['prefix'].'_csrf']))
                $token = $_GET[$this->fDef['prefix'].'_csrf'];

            if (array_key_exists($token, $_SESSION[CSRF_PREFIX.'_csrf'])) {
                $deadline = $_SESSION[CSRF_PREFIX.'_csrf'][$token];
                unset($_SESSION[CSRF_PREFIX.'_csrf'][$token]);
                list($mSec, $sec) = explode(" ", microtime());
                $now = (float)$sec + (float)$mSec;
                if ((float)$deadline >= $now) return TRUE;
            }

            return FALSE;
        }
    }

*/

func (form *Form) ValidateField(k string, value string) bool {
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

func (form *Form) ValidateAll(ctx *fasthttp.RequestCtx) (bool, map[string]string, []string) {
	// https://godoc.org/github.com/valyala/fasthttp#RequestCtx.FormValue
	// https://stackoverflow.com/questions/39265978/get-a-request-parameter-key-value-in-fasthttp

	fValues := map[string]string{}
	wrongFields := []string{}

	for _, k := range form.FieldsOrder {
		if k == "_csrf" {

		} else {
			if !ctx.PostArgs().Has(form.Prefix + k) {
				// Parametero non presente.
				fValues[k] = ""
				if !stringInSlice(k, form.DontValidate) {
					wrongFields = append(wrongFields, k)
				}
			} else {
				fValues[k] = strings.Trim(string(ctx.PostArgs().Peek(form.Prefix+k)), " ")
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

func (form *Form) ImportTable(table *orm.Table, tableT *orm.Dictio, extRefs bool, quickList []string, notRequired []string, prefix string) {
	columns := []string{}
	if len(quickList) > 0 {
		columns = append(columns, quickList...)
	} else {
		for _, v := range table.ColumnsOrder {
			columns = append(columns, v)
		}
	}

	form.Name = table.Name
	form.Prefix = prefix
	for _, v := range columns {
		form.FieldsOrder = append(form.FieldsOrder, v)
		form.Fields[v] = &Field{
			Name:      v,
			Title:     tableT.Title[v],
			MinLength: table.Columns[v].MinLength,
			MaxLength: table.Columns[v].MaxLength,
			Checks:    table.Columns[v].Checks,
		}
		form.UIs[v] = &UI{
			Attrs:      make(map[string]string),
			Default:    "",
			Widget:     "",
			WsUrl:      "",
			WsCallback: "",
			Opts:       tableT.Opts[v],
		}

		// default
		/*
		   "0": "Non definito",
		   "1": "NULL",
		   "2": "Stringa vuota",
		   "3": "Come definito",
		   "4": "CURRENT_TIMESTAMP"
		*/
		if table.Columns[v].UcDefault == "3" {
			form.UIs[v].Default = table.Columns[v].Default
		}

		if !stringInSlice(v, notRequired) {
			form.Required = append(form.Required, v)
			form.UIs[v].Attrs["required"] = "required"
		}

		if table.Columns[v].UI_Widget != "" {
			form.UIs[v].Widget = table.Columns[v].UI_Widget
		}
		if table.Columns[v].UI_WsUrl != "" {
			form.UIs[v].WsUrl = table.Columns[v].UI_WsUrl
		}
		if table.Columns[v].UI_WsCallback != "" {
			form.UIs[v].WsCallback = table.Columns[v].UI_WsCallback
		}
	}

	form.AddCsrfField("")
}
