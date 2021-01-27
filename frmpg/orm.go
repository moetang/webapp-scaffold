package frmpg

import (
	"context"
	"errors"
	"reflect"

	"github.com/moetang/webapp-scaffold/utils"

	"github.com/jackc/pgtype/pgxtype"
)

var ErrNoRecordFound = errors.New("no record found by QuerySingle")

// allowed parameter type: *[]struct{} or *[]*struct{}
func QueryMulti(db pgxtype.Querier, result interface{}, ctx context.Context, sql string, params ...interface{}) error {
	if !checkOrmQueryMultiParamType(result) {
		return errors.New("result type is not *[]struct{} or *[]*struct{}")
	}

	t := getOrmQueryMultiParamItemType(result)
	fm := utils.GetTagStructMapping(t, "mx.orm")

	rs, err := db.Query(
		ctx, sql, params...,
	)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	descs := rs.FieldDescriptions()

	for rs.Next() {
		obj := reflect.New(t)

		var params []interface{}
		for _, v := range descs {
			idx, ok := fm[string(v.Name)]
			if ok {
				params = append(params, obj.Elem().Field(idx).Addr().Interface())
			} else {
				params = append(params, nil)
			}
		}
		err = rs.Scan(params...)
		if err != nil {
			panic(err)
		}

		vr := reflect.ValueOf(result).Elem()
		nvr := reflect.Append(vr, obj)
		vr.Set(nvr)
	}

	return nil
}

// allowed parameter type: *struct{}
func QuerySingle(db pgxtype.Querier, result interface{}, ctx context.Context, sql string, params ...interface{}) error {
	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return errors.New("result type is not *struct{}")
	}
	t = t.Elem()

	vv := reflect.ValueOf(result)
	if vv.IsZero() {
		return errors.New("result should not be nil")
	}

	fm := utils.GetTagStructMapping(t, "mx.orm")

	rs, err := db.Query(
		ctx, sql, params...,
	)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	descs := rs.FieldDescriptions()

	if rs.Next() {
		var params []interface{}
		for _, v := range descs {
			idx, ok := fm[string(v.Name)]
			if ok {
				params = append(params, vv.Elem().Field(idx).Addr().Interface())
			} else {
				params = append(params, nil)
			}
		}
		err = rs.Scan(params...)
		if err != nil {
			panic(err)
		}
		return nil
	}

	return ErrNoRecordFound
}
