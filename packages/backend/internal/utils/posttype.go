package utils

import (
	"errors"
	"slices"
	"strings"
)

func ValidatePostType(postType string) error {
	validPostTypes := []string{"fact", "saying"}
	normalizedPostType := strings.ToLower(postType)

	if !slices.Contains(validPostTypes, normalizedPostType) {
		return errors.New("invalid post type")
	}

	return nil
}
