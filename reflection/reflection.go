package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	v := value(x)

	n := 0
	var field func(int) reflect.Value

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		n = v.NumField()
		field = v.Field
	case reflect.Slice, reflect.Array:
		n = v.Len()
		field = v.Index
	}

	for i := 0; i < n; i++ {
		walk(field(i).Interface(), fn)
	}
}

func value(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}
