package TFilter

import (
	"reflect"
)

func (tf *TFilter) EQ(key string, value interface{}) *TFilter {
	auxObj := []interface{}{}
	maxChan := tf.size / maxChannels(tf.size)

	resChan := make(chan []interface{}, maxChan)
	for i := 0; i < tf.size; i += maxChan {
		go searchEQ(tf.objs[i:i+maxChan], key, value, resChan)
	}

	for i := 0; i < tf.size; i += maxChan {
		select {
		case objs := <-resChan:
			auxObj = append(auxObj, objs...)
		}
	}
	tf.objs = auxObj
	return tf
}

func searchEQ(objs []interface{}, key string, value interface{}, resChan chan []interface{}) {
	auxObj := []interface{}{}
	for _, obj := range objs {
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
	resChan <- auxObj
}
