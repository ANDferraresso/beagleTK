package validator

import (
    "regexp"
)

//
func IsEmail(value string, pars *[]string) bool {
    re := regexp.MustCompile("^[-0-9a-zA-Z.+_]+@[-0-9a-zA-Z.+_]+\\.[a-zA-Z]{2,}$")
    matched := re.Match([]byte(value))
    if matched { return true }
    return false
}

//
func IsIPV4(value string, pars *[]string) bool {
    re := regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")
    matched := re.Match([]byte(value))
    if matched { return true }
    return false
}

//
func IsURL(value string, pars *[]string) bool {
    re := regexp.MustCompile("^(https?:\\/\\/)?([\\da-zA-Z\\.-]+)\\.([a-zA-Z\\.]{2,6})([\\/\\w\\.-]*)*\\/?$")
    matched := re.Match([]byte(value))
    if matched { return true }
    return false
}