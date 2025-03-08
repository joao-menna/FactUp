package utils

import (
	"errors"
	"strconv"
)

func ParseQueryId(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return 0, err
	}

	if id < 1 {
		return 0, errors.New("invalid post id")
	}

	return id, nil
}
