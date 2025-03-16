package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostOffsetLimitBelowOneAndPageBelowZero(t *testing.T) {
	offset := utils.GetPostOffset(0, -1)
	require.Equal(t, 0, offset)
}

func TestPostOffsetLimitOver30(t *testing.T) {
	offset := utils.GetPostOffset(31, 1)
	require.Equal(t, 30, offset)
}

func TestPostOffsetValidInputOutput(t *testing.T) {
	offset := utils.GetPostOffset(20, 3)
	require.Equal(t, 60, offset)
}
