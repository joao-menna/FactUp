package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostMaxCountShouldNotThrowWithZero(t *testing.T) {
	err := utils.CheckPostMaxCountByDay(0)
	require.NoError(t, err)
}

func TestPostMaxCountShouldNotThrowWithTwo(t *testing.T) {
	err := utils.CheckPostMaxCountByDay(2)
	require.NoError(t, err)
}

func TestPostMaxCountShouldThrowWithThree(t *testing.T) {
	err := utils.CheckPostMaxCountByDay(3)
	require.Error(t, err)
}
