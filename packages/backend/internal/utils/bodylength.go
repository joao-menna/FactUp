package utils

import (
	"errors"
	"unicode/utf8"
)

func CheckBodyMaxLength(body string) error {
	if utf8.RuneCountInString(body) > 280 {
		return errors.New("body max length reached")
	}

	return nil
}

func CheckSourceMaxLength(source string) error {
	if utf8.RuneCountInString(source) > 80 {
		return errors.New("source max length reached")
	}

	return nil
}
