package utils

import "regexp"

var phoneRegex string = `^09(1[0-9]|3[1-9]|2[1-9]|9[1-9]|0[1-9]|4[1-9])-?[0-9]{3}-?[0-9]{4}$`

func ValidatePhoneNumber(phone string) bool {
    r, _ := regexp.Compile(phoneRegex)
    return r.MatchString(phone)
}