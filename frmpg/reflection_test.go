package frmpg

import (
	"reflect"
	"testing"

	"github.com/moetang/webapp-scaffold/utils"
)

func TestCheckOrmQueryMultiParamType(t *testing.T) {
	var arr1 []interface{}
	var arr2 []struct{}
	var arr3 []*struct{}
	var arr4 *[]interface{}
	var arr5 *[]struct{}
	var arr6 *[]*struct{}

	utils.AssertFalse(t, checkOrmQueryMultiParamType(arr1))
	utils.AssertFalse(t, checkOrmQueryMultiParamType(arr2))
	utils.AssertFalse(t, checkOrmQueryMultiParamType(arr3))
	utils.AssertFalse(t, checkOrmQueryMultiParamType(arr4))
	utils.AssertTrue(t, checkOrmQueryMultiParamType(arr5))
	utils.AssertTrue(t, checkOrmQueryMultiParamType(arr6))
}

func TestGetOrmQueryMultiParamType(t *testing.T) {
	var arr5 *[]struct{}
	var arr6 *[]*struct{}

	utils.AssertEquals(t, getOrmQueryMultiParamItemType(arr5), reflect.TypeOf(struct{}{}))
	utils.AssertEquals(t, getOrmQueryMultiParamItemType(arr6), reflect.TypeOf(struct{}{}))
}
