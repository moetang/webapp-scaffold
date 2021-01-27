package utils

import (
	"reflect"
	"testing"
)

func TestCheckSlice(t *testing.T) {
	var arr1 []interface{}
	var arr2 []struct{}
	var arr3 []*struct{}
	var arr4 *[]interface{}

	AssertTrue(t, CheckSliceByInstance(arr1))
	AssertTrue(t, CheckSliceByInstance(arr2))
	AssertTrue(t, CheckSliceByInstance(arr3))
	AssertFalse(t, CheckSliceByInstance(arr4))
}

func TestCheckPointer(t *testing.T) {
	var arr1 []interface{}
	var arr2 []struct{}
	var arr3 []*struct{}
	var arr4 *[]interface{}

	AssertFalse(t, CheckPointerByInstance(arr1))
	AssertFalse(t, CheckPointerByInstance(arr2))
	AssertFalse(t, CheckPointerByInstance(arr3))
	AssertTrue(t, CheckPointerByInstance(arr4))
}

func TestGetPointedTypeByInstance(t *testing.T) {
	var arr1 *[]interface{}
	var arr2 *[]struct{}
	var arr3 *[]*struct{}

	AssertEquals(t, GetPointedTypeByInstance(arr1), reflect.TypeOf([]interface{}{}))
	AssertEquals(t, GetPointedTypeByInstance(arr2), reflect.TypeOf([]struct{}{}))
	AssertEquals(t, GetPointedTypeByInstance(arr3), reflect.TypeOf([]*struct{}{}))
}

func TestGetSliceItemTypeByType(t *testing.T) {
	var arr2 []struct{}
	var arr3 []*struct{}

	AssertEquals(t, GetPointedTypeByInstance(arr2), reflect.TypeOf(struct{}{}))
	AssertEquals(t, GetPointedTypeByInstance(arr3), reflect.TypeOf(new(struct{})))
}

func TestGetTagStructMapping(t *testing.T) {
	var arr struct {
		A string `scaff.db:"a"`
		B string `scaff.db:"b"`
		C string `scaff.db:"c"`
	}

	t.Log(GetTagStructMapping(reflect.TypeOf(arr), "scaff.db"))
}
