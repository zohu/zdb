package zdb

import "reflect"

type Value interface {
	string | int | int32 | int64 | float32 | float64 | bool
}

func firstTruthValue[T Value](args ...T) T {
	for _, item := range args {
		if !reflect.ValueOf(item).IsZero() {
			return item
		}
	}
	return args[0]
}
