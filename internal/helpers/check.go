package helpers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// CheckErr checks for x & y error equality during a test.
func CheckErr(t *testing.T, x error, y error, opts ...cmp.Option) {
	if y != nil || x != nil {
		if diff := cmp.Diff(x, y, EquateErrorContent()); diff != "" {
			t.Errorf("error do not match: (-expected +got)\n%s", diff)
		}
		return
	}
}

// Check checks for x & y equality during a test.
func Check(t *testing.T, x interface{}, y interface{}, opts ...cmp.Option) {
	if diff := cmp.Diff(x, y, opts...); diff != "" {
		t.Errorf("do not match: (-expected +got)\n%s", diff)
	}
}
