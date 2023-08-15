package utils

import (
	"regexp"
)

func IsNotNil(value ...interface{}) bool {
	for _, value := range value {
		if value == nil || value == "" {
			return false
		}
	}
	return true
}

func IsUsername(value string) bool {
	if matched, err := regexp.MatchString("[_\\w]{8,64}", value); !matched || err != nil {
		return false
	}
	return true
}

func IsPassword(value string) bool {
	if matched, err := regexp.MatchString("[a-z]{1}", value); !matched || err != nil {
		return false
	}
	if matched, err := regexp.MatchString("[A-Z]{1}", value); !matched || err != nil {
		return false
	}
	if matched, err := regexp.MatchString("[\\d]{1}", value); !matched || err != nil {
		return false
	}
	if matched, err := regexp.MatchString("[_\\W]{1}", value); !matched || err != nil {
		return false
	}
	if matched, err := regexp.MatchString(".{8,64}", value); !matched || err != nil {
		return false
	}
	return true
}

func IsFullName(value string) bool {
	regex, err := regexp.Compile(`^[\w]{1,255}$`)
	if err != nil {
		return false
	}
	return regex.Match([]byte(value))
}
