package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePageShouldThrowWhenParseNegativeNumber(t *testing.T) {
	_, err := utils.ParsePage("-1")
	require.Error(t, err)
}

func TestParsePageShouldNotThrowWhenNumberValid(t *testing.T) {
	page, err := utils.ParsePage("0")
	require.NoError(t, err)
	require.Equal(t, 0, page)

	page, err = utils.ParsePage("1")
	require.NoError(t, err)
	require.Equal(t, 1, page)
}
