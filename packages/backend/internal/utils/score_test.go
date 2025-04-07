package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateScore(t *testing.T) {
	tests := []struct {
		score    int
		expected error
	}{
		{-1, nil},                         // valid score
		{1, nil},                          // valid score
		{-2, errors.New("invalid score")}, // invalid score
		{0, errors.New("invalid score")},  // invalid score
		{2, errors.New("invalid score")},  // invalid score
	}

	for _, test := range tests {
		result := ValidateScore(test.score)
		if test.expected != nil {
			require.EqualError(t, result, test.expected.Error())
		} else {
			require.NoError(t, result)
		}
	}
}
