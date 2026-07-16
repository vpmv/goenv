package env

import (
	"os"
	"strconv"
)

type BasicType interface {
	~string | ~int | ~float64 | ~uint64 | ~bool
}

// Set will convert all basic types to string and set the environment variable
func Set[T BasicType](key string, value T) {
	var s string

	switch any(value).(type) {
	case string:
		s = any(value).(string)
	case int:
		s = strconv.Itoa(any(value).(int))
	case float64:
		s = strconv.FormatFloat(any(value).(float64), 'f', -1, 64) // no trailing zeros
	case uint64:
		s = strconv.FormatUint(any(value).(uint64), 10)
	case bool:
		s = strconv.FormatBool(any(value).(bool))
	default:
		// Should be unreachable due to the constraint
		s = strconv.FormatUint(0, 10)
	}

	_ = os.Setenv(key, s)
}
