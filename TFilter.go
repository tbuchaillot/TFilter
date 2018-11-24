package TFilter

import (
	"reflect"
)

type TFilter struct {
	objs []interface{}
}

func Init(objs interface{}) *TFilter {
	s := reflect.ValueOf(objs)
	if s.Kind() != reflect.Slice {
		panic("given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	tf := new(TFilter)
	tf.objs = ret

	return tf
}

func (tf *TFilter) GetObjs() []interface{} {
	return tf.objs
}
