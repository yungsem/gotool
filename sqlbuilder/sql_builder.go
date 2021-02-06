package sqlbuilder

import (
	reflectx "github.com/yungsem/gotool/relectx"
	"strings"
)

type SqlBuilder struct {
}

func (b *SqlBuilder) buildInsert(i interface{}) (string, error) {
	tblName, err := reflectx.GetStructNameSnakeUpper(i)
	if err != nil {
		return "", err
	}

	fieldNames, err := reflectx.GetFiledNamesSnakeUpper(i)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(tblName)
	sb.WriteString(" (")
	sb.WriteString(strings.Join(fieldNames, ","))
	sb.WriteString(")")
	sb.WriteString(" VALUES")
	sb.WriteString(" (")
	for i := 0; i < len(fieldNames); i++ {
		sb.WriteString("?")
		if i != len(fieldNames)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")

	return sb.String(), nil
}

func (b *SqlBuilder) buildUpdateById(i interface{}) (string, error) {
	tblName, err := reflectx.GetStructNameSnakeUpper(i)
	if err != nil {
		return "", err
	}

	fieldNames, err := reflectx.GetFiledNamesSnakeUpper(i)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("UPDATE ")
	sb.WriteString(tblName)
	sb.WriteString(" SET ")
	for i, v := range fieldNames {
		if v == "ID" || v == "CREATE_TIME" || v == "CREATE_USER" {
			continue
		}
		sb.WriteString(v)
		sb.WriteString("=?")
		if i != len(fieldNames)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(" WHERE ID=?")

	return sb.String(), nil
}

func (b *SqlBuilder) buildGetById(i interface{}) (string, error) {
	tblName, err := reflectx.GetStructNameSnakeUpper(i)
	if err != nil {
		return "", err
	}

	fieldNames, err := reflectx.GetFiledNamesSnakeUpper(i)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(fieldNames, ","))
	sb.WriteString(" FROM ")
	sb.WriteString(tblName)
	sb.WriteString(" WHERE ID=?")

	return sb.String(), nil
}
