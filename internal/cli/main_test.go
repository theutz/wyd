package cli

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name   string
		wants  any
		errMsg string
	}{
		{name: "zero value", wants: nil, errMsg: "unexpected error"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := New()
			assert.Zero(t, got)
			assert.EqualError(t, err, tc.errMsg)
		})
	}
}
