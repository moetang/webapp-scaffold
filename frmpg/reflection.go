package frmpg

import (
	"reflect"

	"github.com/moetang/webapp-scaffold/utils"
)

func checkOrmQueryMultiParamType(i interface{}) bool {
	if !utils.CheckPointerByInstance(i) {
		return false
	}
	rt := utils.GetPointedTypeByInstance(i)
	if !utils.CheckSliceByType(rt) {
		return false
	}
	et := utils.GetSliceItemTypeByType(rt)
	if utils.CheckPointerByType(et) {
		et = et.Elem()
	}
	if et.Kind() != reflect.Struct {
		return false
	}
	return true
}

func getOrmQueryMultiParamItemType(i interface{}) reflect.Type {
	rt := utils.GetPointedTypeByInstance(i)
	et := utils.GetSliceItemTypeByType(rt)
	if utils.CheckPointerByType(et) {
		return et.Elem()
	} else {
		return et
	}
}
