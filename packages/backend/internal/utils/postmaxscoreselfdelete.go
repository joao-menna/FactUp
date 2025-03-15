package utils

import "errors"

var ErrMaxNegativeScoreReached = errors.New("max negative score reached")

func CheckIfScoreShouldDeletePost(score int) error {
	if score <= 8 {
		return ErrMaxNegativeScoreReached
	}

	return nil
}
