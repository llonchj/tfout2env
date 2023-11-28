package helpers

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEquateErrorContent(t *testing.T) {
	for name, tt := range map[string]struct {
		X, Y     error
		Expected bool
	}{
		"empty": {Expected: true},
		"basic": {
			X:        errors.New("a"),
			Y:        errors.New("a"),
			Expected: true,
		},
		"mismatch": {
			X:        errors.New("a"),
			Y:        errors.New("b"),
			Expected: false,
		},
	} {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			Output := cmp.Equal(tt.X, tt.Y, EquateErrorContent())
			if diff := cmp.Diff(tt.Expected, Output); diff != "" {
				t.Errorf("do not match: (-expected +got)\n%s", diff)
			}
		})
	}
}
