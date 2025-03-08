package utils

import (
	"errors"
	"strconv"
)

func ParsePostLimit(limit string) (int, error) {
	l, err := strconv.Atoi(limit)

	if err != nil {
		return 0, err
	}

	if l < 1 || l > 30 {
		return 0, errors.New("invalid limit")
	}

	return l, nil
}
