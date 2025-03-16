package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNegativePostLimitValue(t *testing.T) {
	_, err := utils.ParsePostLimit("-1")
	require.Error(t, err)
}

func TestZeroPostLimitValue(t *testing.T) {
	_, err := utils.ParsePostLimit("0")
	require.Error(t, err)
}

func TestTextStringPostLimitValue(t *testing.T) {
	_, err := utils.ParsePostLimit("abc")
	require.Error(t, err)
}

func TestOver30PostLimitValue(t *testing.T) {
	_, err := utils.ParsePostLimit("31")
	require.Error(t, err)
}

func TestInRangePostLimitValue(t *testing.T) {
	id, err := utils.ParsePostLimit("1")
	require.NoError(t, err)
	require.Equal(t, 1, id)

	id, err = utils.ParsePostLimit("30")
	require.NoError(t, err)
	require.Equal(t, 30, id)
}
