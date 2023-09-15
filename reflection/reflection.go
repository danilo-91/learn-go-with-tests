package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	v := value(x)
	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walk(v.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}
	}
}

func value(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}
