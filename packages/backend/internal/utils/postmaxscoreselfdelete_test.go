package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckIfScoreShouldDeletePostShouldNotThrowWithZero(t *testing.T) {
	err := utils.CheckIfScoreShouldDeletePost(0)
	require.NoError(t, err)
}

func TestCheckIfScoreShouldDeletePostShouldNotThrowWithSeven(t *testing.T) {
	err := utils.CheckIfScoreShouldDeletePost(-7)
	require.NoError(t, err)
}

func TestCheckIfScoreShouldDeletePostShouldThrowWithEight(t *testing.T) {
	err := utils.CheckIfScoreShouldDeletePost(-8)
	require.Error(t, err)
}
