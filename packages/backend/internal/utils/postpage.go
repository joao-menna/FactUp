package utils

import (
	"errors"
	"strconv"
)

func ParsePage(pageStr string) (int, error) {
	p, err := strconv.Atoi(pageStr)

	if err != nil {
		return 0, err
	}

	if p < 0 {
		return 0, errors.New("invalid page")
	}

	return p, nil
}
