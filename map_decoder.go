package fakers

import (
	"fmt"
	"reflect"

	"github.com/vingarcia/structscanner"
)

// MapTagDecoder can be used to fill a struct with the values of a map.
//
// It works recursively so you can pass nested structs to it.
type customValuesMapDecoder struct {
	customValues map[string]any
}

// customValuesMapDecoder returns a new decoder for filling a given struct.
// It will try to find values for the fields by looking at the provided
// map of customValues (indexing by the name of attribute), if the specific
// attribute doesn't have a custom value a default deterministic value
// is generated instead.
func newCustomValuesMapDecoder(customValues map[string]any) customValuesMapDecoder {
	return customValuesMapDecoder{
		customValues: customValues,
	}
}

// DecodeField implements the TagDecoder interface
func (c customValuesMapDecoder) DecodeField(info structscanner.Field) (any, error) {
	if c.customValues[info.Name] != nil {
		return c.customValues[info.Name], nil
	}

	switch info.Kind {

	case reflect.String:
		return "Fake" + info.Name, nil

	case reflect.Bool:
		return false, nil

	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return 42, nil

	case reflect.Float32, reflect.Float64:
		return 42.0, nil

	case reflect.Uintptr:
		return 0, nil

	case
		reflect.Map, reflect.Slice, reflect.Interface,
		reflect.Pointer, reflect.Func, reflect.Chan:
		return nil, nil

	case reflect.Array:
		// This should cause the array to be unnafected,
		// which is the same as initializing it with zero
		// on all fields:
		return nil, nil

	case reflect.Struct:
		return newCustomValuesMapDecoder(nil), nil

	default:
		// The list of currently unsupported "kinds" is:
		// reflect.Complex64
		// reflect.Complex128
		// reflect.UnsafePointer
		//
		// TODO(vgarcia): learn more about these types and
		// consider if we should add support

		return nil, fmt.Errorf("unsupported type for faker: %v", info.Type)
	}
}
