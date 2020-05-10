package application

import (
	"fmt"
	"log"
	"os"
	"reflect"
)


var (
	NO_ERROR_MESSAGE_PASSED="No Error Message exists"
)

// Check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Error Logging
func error_logging(msg string)  {
	if len(msg) == 0 {
		Red(NO_ERROR_MESSAGE_PASSED)
		os.Exit(1)
	}
	Red(msg)
	os.Exit(1)
}

// Fatal Error
func _fatal_error(err error)  {
	log.Fatalf("error: %v", err)
	os.Exit(1)
}

func _ctos(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func isZero(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool() == false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0

	case reflect.Ptr, reflect.Interface:
		return isZero(v.Elem())

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.String, reflect.Map:
		return v.Len() == 0

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	// reflect.Chan, reflect.UnsafePointer, reflect.Func
	default:
		return v.IsNil()
	}
}

// IsZero reports whether v is zero struct
// Does not support cycle pointers for performance, so as json
func IsZero(v interface{}) bool {
	return isZero(reflect.ValueOf(v))
}