package shared

import "reflect"

// IsZero returns whether the value is the natural zero value
func IsZero(x interface{}) bool {
	return x == nil || x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
