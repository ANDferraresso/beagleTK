package form

import (
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