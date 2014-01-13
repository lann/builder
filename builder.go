package builder

import (
	"go/ast"
	"reflect"
	"github.com/mndrix/ps"
)

type Builder struct {
	builderMap ps.Map
}

var emptyBuilder = Builder{ps.NewMap()}

type any interface{}

func reflectConvert(val any, toType reflect.Type) any {
	return reflect.ValueOf(val).Convert(toType).Interface()
}

func getBuilderMap(builder any) ps.Map {
	b := reflectConvert(builder, reflect.TypeOf(Builder{})).(Builder)

	if b.builderMap == nil {
		return ps.NewMap()
	}

	return b.builderMap
}

func Set(builder any, name string, val any) any {
	b := Builder{getBuilderMap(builder).Set(name, val)}
	return reflectConvert(b, reflect.TypeOf(builder))
}

func Append(builder any, name string, vals ...any) any {
	maybeList, ok := getBuilderMap(builder).Lookup(name)

	var list ps.List
	if ok {
		list, ok = maybeList.(ps.List)
	}
	if !ok {
		list = ps.NewList()
	}

	for _, val := range vals {
		list = list.Cons(val)
	}
	return Set(builder, name, list)
}

func listToSlice(list ps.List, arrayType reflect.Type) reflect.Value {
	size := list.Size()
	slice := reflect.MakeSlice(arrayType, size, size)
	for i := size - 1; i >= 0; i--  {
		val := reflect.ValueOf(list.Head())
		slice.Index(i).Set(val)
		list = list.Tail()
	}
	return slice
}

var anyArrayType = reflect.TypeOf([]interface{}{})

func Get(builder any, name string) (any, bool) {
	val, ok := getBuilderMap(builder).Lookup(name)
	if !ok {
		return nil, false
	}

	list, isList := val.(ps.List)
	if isList {
		arrayType := anyArrayType

		if ast.IsExported(name) {
			structType := getBuilderStructType(reflect.TypeOf(builder))
			if structType != nil {
				field, ok := (*structType).FieldByName(name)
				if ok {
					arrayType = field.Type
				}
			}
		}

		val = listToSlice(list, arrayType).Interface()
	}

	return val, true
}

func GetStruct(builder any) any {
	structVal := newBuilderStruct(reflect.TypeOf(builder))
	if structVal == nil {
		return nil
	}

	getBuilderMap(builder).ForEach(func(name string, val ps.Any) {
		if ast.IsExported(name) {
			field := structVal.FieldByName(name)

			list, isList := val.(ps.List)
			if isList {
				val = listToSlice(list, field.Type()).Interface()
			}

			field.Set(reflect.ValueOf(val))
		}
	})

	return structVal.Interface()
}
