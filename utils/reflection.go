package utils

import (
	"errors"
	"reflect"
)

func CheckSliceByInstance(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Slice {
		return true
	} else {
		return false
	}
}

func CheckSliceByType(t reflect.Type) bool {
	if t.Kind() == reflect.Slice {
		return true
	} else {
		return false
	}
}

func CheckPointerByInstance(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		return true
	} else {
		return false
	}
}

func CheckPointerByType(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		return true
	} else {
		return false
	}
}

func GetPointedTypeByInstance(i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}

func GetSliceItemTypeByType(t reflect.Type) reflect.Type {
	return t.Elem()
}

func GetTagStructMapping(t reflect.Type, tagName string) map[string]int {
	r := make(map[string]int)

	fnum := t.NumField()
	for i := 0; i < fnum; i++ {
		f := t.Field(i)
		k, ok := f.Tag.Lookup(tagName)
		if ok {
			_, pok := r[k]
			if pok {
				panic(errors.New("tag value exists:" + tagName + "=" + k))
			}
			r[k] = i
		}
	}
	return r
}
