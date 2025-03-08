package utils_test

import (
	"backend/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidFactPostType(t *testing.T) {
	postType := "Fact"

	err := utils.ValidatePostType(postType)
	require.NoError(t, err)
}

func TestValidSayingPostType(t *testing.T) {
	postType := "Saying"

	err := utils.ValidatePostType(postType)
	require.NoError(t, err)
}

func TestInvalidPostType(t *testing.T) {
	postType := "abc"

	err := utils.ValidatePostType(postType)
	require.Error(t, err)
}
