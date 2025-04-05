package utils

import "errors"

func CheckBodyMinLength(body string) error {
	if len(body) <= 5 {
		return errors.New("body min length reached")
	}

	return nil
}

func CheckBodyMaxLength(body string) error {
	if len(body) > 280 {
		return errors.New("body max length reached")
	}

	return nil
}

func CheckSourceMaxLength(source string) error {
	if len(source) > 80 {
		return errors.New("source max length reached")
	}

	return nil
}
