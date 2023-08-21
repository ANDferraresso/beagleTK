package validator

import (
)

type Check struct {
    Func string
    Pars []string
}

type Validator struct {
    Funcs map[string]func(value string, pars *[]string) bool
}

//
func (val *Validator) SetupValidator() {
    val.Funcs = map[string]func(value string, pars *[]string) bool {
        // Internet
        "isEmail": IsEmail,
        "isIPV4": IsIPV4,
        "isURL": IsURL,
        // Numeric
        "isDecimal": IsDecimal,
        "isMoney": IsMoney,
        "isInteger": IsInteger,
        "isNegativeInt": IsNegativeInt,
        "isPositiveInt": IsPositiveInt,
        "isZeroNegativeInt": IsZeroNegativeInt,
        "isZeroPositiveInt": IsZeroPositiveInt,
        "isIntGreater": IsIntGreater,
        "isIntGreaterEqual": IsIntGreaterEqual,
        "isIntInRange": IsIntInRange,
        "isIntLower": IsIntLower,
        "isIntLowerEqual": IsIntLowerEqual,
        "isPassword": IsPassword,
        "allowedChars": AllowedChars,
        "forbiddenChars": ForbiddenChars,
        "isStringEqual": IsStringEqual,
        "isLength": IsLength,
        "isLengthInRange": IsLengthInRange,
        "isMaxLength": IsMaxLength,
        "isMinLength": IsMinLength,
        "isRegex": IsRegex,
    }
}

//
func (val *Validator) Validate(f string, value string, pars *[]string) bool {
    return val.Funcs[f](value, pars)
}