package utils

import (
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		data[t.Field(i).Name] = v.Field(i).Interface()

	}

	return data
}

func StructToMapJson(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		jsonKey := t.Field(i).Tag.Get("json")

		if jsonKey != "-" {
			data[jsonKey] = v.Field(i).Interface()
		}

	}

	return data
}
