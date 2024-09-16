package app

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name  string
		given string
		wants *Config
	}{
		{
			name:  "normal db path",
			given: "~/.local/share/wyd/wyd.db",
			wants: &Config{
				DatabasePath: "/Users/michael/.local/share/wyd/wyd.db",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange

			// Act
			c, err := New(tc.given)

			// Assert
			assert.NoError(t, err)
			assert.Equal(
				t,
				tc.wants,
				c.config,
				assert.OmitEmpty(),
			)
		})
	}
}
