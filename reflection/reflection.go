package main

import "reflect"


func walk(x interface{}, fn func(string)) {
    v := reflect.ValueOf(x)
    field := v.Field(0)
    fn(field.String())
}
