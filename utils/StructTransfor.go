package utils

import (
	"errors"
	"fmt"
	"main/Const"
	"reflect"
)

// ToInterfaceSlice 转换为一维泛型
func ToInterfaceSlice(data interface{}) []interface{} {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Struct:
		len := v.NumField()
		result := make([]interface{}, len)
		for i := 0; i < len; i++ {
			result[i] = v.Field(i).Interface()
		}
		return result
	default:
		return []interface{}{data}
	}
}

// OneToDoubleInterface 一维泛型结构转换为二维泛型
func OneToDoubleInterface(data []interface{}) (re [][]interface{}) {
	for i, value := range data {
		v := reflect.ValueOf(value)
		re = append(re, []interface{}{})
		for k := 0; k < v.NumField(); k++ {
			re[i] = append(re[i], v.Field(k).Interface())
		}
	}
	return re
}

func ToDoubleInterfaceSlice(slice interface{}) ([][]interface{}, error) {
	//value获取
	v := reflect.ValueOf(slice)
	//如果
	if v.Kind() != reflect.Slice {
		return [][]interface{}{}, errors.New(Const.ErrorDataType)
		//panic("ToInterfaceSlice() given a non-slice type")
	}

	// Iterate over the slice and convert each element to []interface{}
	result := make([][]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		//单个数据获取
		elem := v.Index(i)
		//数据需要为结构体
		if elem.Kind() != reflect.Struct {
			return [][]interface{}{}, errors.New(Const.ErrorDataType)
			//panic("ToInterfaceSlice() can only handle slices of struct")
		}

		fields := make([]interface{}, elem.NumField())
		ptr := 0 //模拟指针
		for j := 0; j < elem.NumField(); j++ {
			eme := elem.Field(j)
			//嵌套结构体判断
			if eme.Kind() == reflect.Struct {
				for k := 0; k < eme.NumField(); k++ {
					//必须扩充数组
					fields = append(fields, []interface{}{})
					fields[ptr] = eme.Field(k).Interface()
					ptr++
					fmt.Println("struct----------------------")
				}
			} else {
				fields[ptr] = elem.Field(j).Interface()
				ptr++
			}
		}

		result[i] = fields
	}

	return result, nil
}
