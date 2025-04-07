package utils

import "errors"

func ValidateScore(score int) error {
	if score != -1 && score != 1 {
		return errors.New("invalid score")
	}

	return nil
}
