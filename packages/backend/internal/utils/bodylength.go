package utils

import "errors"

func CheckBodyMaxLength(body string) error {
	if len(body) > 280 {
		return errors.New("body max length reached")
	}

	return nil
}

func CheckSourceMaxLength(source string) error {
	if len(source) > 128 {
		return errors.New("source max length reached")
	}

	return nil
}
