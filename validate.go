package goonung

import (
	"reflect"
)

var (
	is  bool
	err error
)

func validate(param interface{}, kinds ...reflect.Kind) (bool, error) {
	var val = reflect.ValueOf(param)

	for _, kind := range kinds {
		if kind == val.Kind() {
			return true, nil
		}
	}

	return false, nil
}

// IsArray : method to check the param is array
func IsArray(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Array)

	return is, err
}

// IsBool : method to check the param is boolean
func IsBool(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Bool)

	return is, err
}

// IsComplex : method to check the param is complex
func IsComplex(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Complex64, reflect.Complex128)

	return is, err
}

// IsFloat : method to check the param is float
func IsFloat(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Float32, reflect.Float64)

	return is, err
}

// IsInt : method to check the param is integer
func IsInt(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64)

	return is, err
}

// IsSlice : method to check the param is slice
func IsSlice(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Slice)

	return is, err
}

// IsString : method to check the param is string
func IsString(param interface{}) (bool, error) {
	is, err = validate(param, reflect.String)

	return is, err
}

// IsStruct : method to check the param is struct
func IsStruct(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Struct)

	return is, err
}

// IsUint : method to check the param is unsigned integer
func IsUint(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64)

	return is, err
}

// IsUintPtr : method to check the param is unsigned integer pointer (unsafe)
func IsUintPtr(param interface{}) (bool, error) {
	is, err = validate(param, reflect.Uintptr)

	return is, err
}
