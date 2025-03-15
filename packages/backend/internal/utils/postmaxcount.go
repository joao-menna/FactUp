package utils

import "errors"

var ErrMaxPostForToday = errors.New("reached maximum post count for this account today")

func CheckPostMaxCountByDay(count int) error {
	if count >= 3 {
		return ErrMaxPostForToday
	}

	return nil
}
