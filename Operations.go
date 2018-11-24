package TFilter

import "reflect"

func (tf *TFilter) EQ(key string, value interface{}) *TFilter {
	auxObj := []interface{}{}
	for _, obj := range tf.objs {
		RValue := reflect.ValueOf(obj)
		RType := reflect.TypeOf(obj)
		for i := 0; i < RValue.NumField(); i++ {
			fieldRType := RType.Field(i)
			fieldRValue := RValue.Field(i)

			zero := reflect.Zero(fieldRValue.Type()).Interface()
			isZero := reflect.DeepEqual(fieldRValue.Interface(), zero)
			if !isZero && fieldRType.Tag.Get("key") == key {
				if fieldRValue.Interface() == value {
					auxObj = append(auxObj, obj)
				}
			}
		}
	}
	tf.objs = auxObj
	return tf

}
