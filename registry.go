package builder

import (
	"reflect"
)

var registry = make(map[reflect.Type]reflect.Type)

func Register(builderProto interface{}, structProto interface{}) any {
	empty := RegisterType(
		reflect.TypeOf(builderProto),
		reflect.TypeOf(structProto),
	).Interface()
	return empty
}

func RegisterType(builderType reflect.Type, structType reflect.Type) *reflect.Value {
	structType.NumField() // Panic if structType is not a struct
	registry[builderType] = structType
	emptyValue := reflect.ValueOf(emptyBuilder).Convert(builderType)
	return &emptyValue
}

func getBuilderStructType(builderType reflect.Type) *reflect.Type {
	structType, ok := registry[builderType]
	if !ok {
		return nil
	}
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
