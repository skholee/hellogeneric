package main

import "reflect"

func FilterByReflect(elems, predicate interface{}) interface{} {
	elemsValue := reflect.ValueOf(elems)

	if elemsValue.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	predicateValue := reflect.ValueOf(predicate)
	if predicateValue.Kind() != reflect.Func {
		panic("filter: wrong type, not a func")
	}

	if (predicateValue.Type().NumIn() != 1) ||
		(predicateValue.Type().NumOut() != 1) ||
		(predicateValue.Type().In(0) != elemsValue.Type().Elem()) ||
		(predicateValue.Type().Out(0) != reflect.TypeOf(true)) {
		panic("filter: wrong type, predicate must be of type func(" +
			elemsValue.Elem().String() + ") bool")
	}

	var indexes []int
	for i := 0; i < elemsValue.Len(); i++ {
		if predicateValue.Call([]reflect.Value{elemsValue.Index(i)})[0].Bool() {
			indexes = append(indexes, i)
		}
	}

	r := reflect.MakeSlice(elemsValue.Type(), len(indexes), len(indexes))
	for i := range indexes {
		r.Index(i).Set(elemsValue.Index(indexes[i]))
	}
	return r.Interface()
}
