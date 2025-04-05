package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckBodyMinLength(t *testing.T) {
	tests := []struct {
		body      string
		expectErr bool
	}{
		{"", true},                      // Length 0
		{"1234", true},                  // Length 4
		{"12345", true},                 // Length 5
		{"123456", false},               // Length 6
		{"This is a valid body", false}, // Length > 5
	}

	for _, test := range tests {
		err := CheckBodyMinLength(test.body)
		if test.expectErr {
			require.Error(t, err, "CheckBodyMinLength should return an error for body: %q", test.body)
		} else {
			require.NoError(t, err, "CheckBodyMinLength should not return an error for body: %q", test.body)
		}
	}
}

func TestCheckBodyMaxLength(t *testing.T) {
	tests := []struct {
		body     string
		expected bool
	}{
		{"short body", false},
		{string(make([]byte, 280)), false},
		{"this body is way too long and exceeds the maximum length of 280 characters. This is just to ensure that we go over the limit and trigger the error that we expect to receive from the CheckBodyMaxLength function. I have to make this longer, because Cline was dumb and made a string with less than 280 characters.", true},
	}

	for _, test := range tests {
		err := CheckBodyMaxLength(test.body)
		if test.expected {
			require.Error(t, err, "CheckBodyMaxLength should return an error for body: %q", test.body)
		} else {
			require.NoError(t, err, "CheckBodyMaxLength should not return an error for body: %q", test.body)
		}
	}
}

func TestCheckSourceMaxLength(t *testing.T) {
	tests := []struct {
		source   string
		expected bool
	}{
		{"short source", false},
		{string(make([]byte, 80)), false},
		{"this source is way too long and exceeds the maximum length of 80 characters. This is just to ensure that we go over the limit and trigger the error that we expect to receive from the CheckSourceMaxLength function.", true},
	}

	for _, test := range tests {
		err := CheckSourceMaxLength(test.source)
		if test.expected {
			require.Error(t, err, "CheckSourceMaxLength should return an error for source: %q", test.source)
		} else {
			require.NoError(t, err, "CheckSourceMaxLength should not return an error for source: %q", test.source)
		}
	}
}
