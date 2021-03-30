package reflectx

import (
	"errors"
	"github.com/yungsem/gotool/strings"
	"reflect"
)

var TypeNotSupportError = errors.New("type must be struct")

func FiledNames(i interface{}) ([]string, error) {
	return fieldNames(i, nil)
}

func FiledNamesSnakeLower(i interface{}) ([]string, error) {
	return fieldNames(i, strings.ToSnakeCaseLower)
}

func FiledNamesSnakeUpper(i interface{}) ([]string, error) {
	return fieldNames(i, strings.ToSnakeCaseUpper)
}

func fieldNames(i interface{}, caseConverter func(string) string) ([]string, error) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, TypeNotSupportError
	}
	var names []string
	for idx := 0; idx < t.NumField(); idx++ {
		field := t.Field(idx)
		st := field.Type
		if st.Kind() != reflect.Struct {
			if caseConverter != nil {
				names = append(names, caseConverter(field.Name))
			} else {
				names = append(names, field.Name)
			}
			continue
		}
		// 如果字段也是 struct ，则需要再处理一层
		// 不递归，只处理一层嵌套
		for sIdx := 0; sIdx < st.NumField(); sIdx++ {
			sField := st.Field(sIdx)
			if caseConverter != nil {
				names = append(names, caseConverter(sField.Name))
			} else {
				names = append(names, sField.Name)
			}
		}
	}
	return names, nil
}

func StructName(i interface{}) (string, error) {
	return structName(i, nil)
}

func StructNameSnakeLower(i interface{}) (string, error) {
	return structName(i, strings.ToSnakeCaseLower)
}

func StructNameSnakeUpper(i interface{}) (string, error) {
	return structName(i, strings.ToSnakeCaseUpper)
}

// structName 返回 i 的 struct name ，i 必须是 struct 类型
// caseConverter 是一个转换器，用于将获取到的 struct name 转换为相应的格式
func structName(i interface{}, caseConverter func(string) string) (string, error) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", TypeNotSupportError
	}

	if caseConverter != nil {
		return caseConverter(t.Name()), nil
	}
	return t.Name(), nil
}
