package reflectx

import (
	"errors"
	"github.com/yungsem/gotool/strings"
	"reflect"
)

var TypeNotSupportError = errors.New("type must be struct")

func GetFiledNames(i interface{}) ([]string, error) {
	return getFieldNames(i, nil)
}

func GetFiledNamesSnakeLower(i interface{}) ([]string, error) {
	return getFieldNames(i, strings.ToSnakeCaseLower)
}

func GetFiledNamesSnakeUpper(i interface{}) ([]string, error) {
	return getFieldNames(i, strings.ToSnakeCaseUpper)
}

func getFieldNames(i interface{}, caseConverter func(string) string) ([]string, error) {
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
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			for sIdx := 0; sIdx < fieldType.NumField(); sIdx++ {
				sField := fieldType.Field(sIdx)
				sFieldName := sField.Name
				if caseConverter != nil {
					sFieldName = caseConverter(sFieldName)
				}
				names = append(names, sFieldName)
			}
		} else {
			fieldName := field.Name
			if caseConverter != nil {
				fieldName = caseConverter(fieldName)
			}
			names = append(names, fieldName)
		}
	}
	return names, nil
}

func GetStructName(i interface{}) (string, error) {
	return getStructName(i, nil)
}

func GetStructNameSnakeLower(i interface{}) (string, error) {
	return getStructName(i, strings.ToSnakeCaseLower)
}

func GetStructNameSnakeUpper(i interface{}) (string, error) {
	return getStructName(i, strings.ToSnakeCaseUpper)
}

func getStructName(i interface{}, caseConverter func(string) string) (string, error) {
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
