package validator

import (
    "regexp"
    "strconv"
)

//
func IsDecimal(value string, pars *[]string) bool {
    re := regexp.MustCompile("(^\\-{0,1}0{1}\\.[0-9]+$)|(^\\-{0,1}[1-9]+[0-9]*\\.[0-9]+$)") // -1.0, 10.52, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

//
func IsMoney(value string, pars *[]string) bool {
    re := regexp.MustCompile("(^0{1}\\.[0-9]{2}$)|(^[1-9]+[0-9]*\\.[0-9]{2}$)") // 0.00, 1.20, 19.50, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

//
func IsInteger(value string, pars *[]string) bool {
    re := regexp.MustCompile("^\\-[1-9]+[0-9]*|[0-9]+[0-9]*$") // 0, 1, -2, -123, 456, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}


// 
func IsIntInRange(value string, pars *[]string) bool {
	var v int64
	var a int64
    var b int64
    var err error
	v, err = strconv.ParseInt(*value, 10, 64)
    if err != nil { 
        return false 
    }
    a, err = strconv.ParseInt((*pars)[0], 10, 64)
    if err != nil { 
        return false 
    }
    b, err = strconv.ParseInt((*pars)[1], 10, 64)
    if v >= a && v <= b {
        return true
    }
    return false

//
func IsNegativeInt(value string, pars *[]string) bool {
    re := regexp.MustCompile("^\\-[1-9]+[0-9]*$") // -1, -10, -123, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

//
func IsPositiveInt(value string, pars *[]string) bool {
    re := regexp.MustCompile("^[1-9]+[0-9]*$") // 1, 10, 123, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

//
func IsZeroNegativeInt(value string, pars *[]string) bool {
    re := regexp.MustCompile("(^[0]{1}$)|(^\\-[1-9]+[0-9]*$)") // 0, -1, -10, -123, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

//
func IsZeroPositiveInt(value string, pars *[]string) bool {
    re := regexp.MustCompile("(^[0]{1}$)|(^[1-9]+[0-9]*$)") // 0, 1, 10, 123, ...
    matched := re.Match([]byte(value))
    if matched { 
        return true 
    }
    return false
}

// COMPARING


//
func IsIntGreater(value string, pars *[]string) bool {
    x, err := strconv.ParseInt(value, 10, 64)
    if err != nil { 
        return false 
    }
    p, err := strconv.ParseInt((*pars)[0], 10, 64)
    if x > p { 
        return true 
    }
    return false
}

//
func IsIntGreaterEqual(value string, pars *[]string) bool {
    x, err := strconv.ParseInt(value, 10, 64)
    if err != nil { 
        return false 
    }
    p, err := strconv.ParseInt((*pars)[0], 10, 64)
    if x >= p { 
        return true 
    }
    return false
}

//
func IsIntInRange(value string, pars *[]string) bool {
    x, err := strconv.ParseInt(value, 10, 64)
    if err != nil { 
        return false 
    }
    p1, err := strconv.ParseInt((*pars)[0], 10, 64)
    if err != nil { 
        return false 
    }
    p2, err := strconv.ParseInt((*pars)[1], 10, 64)
    if err != nil { 
        return false 
    }
    if x >= p1 && x <= p2 { 
        return true 
    }
    return false
}

//
func IsIntLower(value string, pars *[]string) bool {
    x, err := strconv.ParseInt(value, 10, 64)
    if err != nil { 
        return false 
    }
    p, err := strconv.ParseInt((*pars)[0], 10, 64)
    if x >= p { 
        return true 
    }
    if x < p { 
        return true 
    }
    return false
}

//
func IsIntLowerEqual(value string, pars *[]string) bool {
    x, err := strconv.ParseInt(value, 10, 64)
    if err != nil { 
        return false 
    }
    p, err := strconv.ParseInt((*pars)[0], 10, 64)
    if x >= p { 
        return true 
    }
    if x <= p { 
        return true 
    }
    return false
}