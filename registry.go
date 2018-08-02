package builder

import (
	"reflect"
	"sync"
)

var registry sync.Map

// RegisterType maps the given builderType to a structType.
// This mapping affects the type of slices returned by Get and is required for
// GetStruct to work.
//
// Returns a Value containing an empty instance of the registered builderType.
//
// RegisterType will panic if builderType's underlying type is not Builder or
// if structType's Kind is not Struct.
func RegisterType(builderType, structType reflect.Type) *reflect.Value {
	structType.NumField() // Panic if structType is not a struct
	registry.Store(builderType, structType)
	emptyValue := emptyBuilderValue.Convert(builderType)
	return &emptyValue
}

// Register wraps RegisterType, taking instances instead of Types.
//
// Returns an empty instance of the registered builder type which can be used
// as the initial value for builder expressions. See example.
func Register(builderProto interface{}, structProto interface{}) interface{} {
	empty := RegisterType(
		reflect.TypeOf(builderProto),
		reflect.TypeOf(structProto),
	).Interface()
	return empty
}

func getBuilderStructType(builderType reflect.Type) *reflect.Type {
	t, ok := registry.Load(builderType)
	if !ok {
		return nil
	}

	structType := t.(reflect.Type)
	return &structType
}

func newBuilderStruct(builderType reflect.Type) *reflect.Value {
	structType := getBuilderStructType(builderType)
	if structType == nil {
		return nil
	}
	newStruct := reflect.New(*structType).Elem()
	return &newStruct
}
