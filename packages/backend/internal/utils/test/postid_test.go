package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNegativePostIdValue(t *testing.T) {
	_, err := utils.ParsePostId("-1")

	require.Error(t, err)
}

func TestZeroPostIdValue(t *testing.T) {
	_, err := utils.ParsePostId("0")

	require.Error(t, err)
}

func TestValidPostIdValue(t *testing.T) {
	id, err := utils.ParsePostId("1")

	require.NoError(t, err)
	require.Equal(t, 1, id)
}
