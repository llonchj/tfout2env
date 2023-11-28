package helpers

import (
	"github.com/google/go-cmp/cmp"
)

func EquateErrorContent() cmp.Option {
	return cmp.FilterValues(
		func(x, y interface{}) bool {
			_, ok1 := x.(error)
			_, ok2 := y.(error)
			return ok1 && ok2
		},
		cmp.Comparer(func(x, y interface{}) bool {
			if x == nil || y == nil {
				return x == nil && y == nil
			}
			xe := x.(error)
			ye := y.(error)
			return xe.Error() == ye.Error()
		}),
	)
}
