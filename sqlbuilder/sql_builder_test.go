package sqlbuilder

import (
	"testing"
	"time"
)

type BaseModel struct {
	ID         int
	CreateTime time.Time
	CreateUser string
	UpdateTime time.Time
	UpdateUser string
	Deleted    int
}

type User struct {
	BaseModel
	Username string
	Password string
}

func TestBuilderInsert(t *testing.T) {
	var builder SqlBuilder
	u := User{
		BaseModel: BaseModel{
			ID:         1,
			CreateTime: time.Now(),
			CreateUser: "",
			UpdateTime: time.Now(),
			UpdateUser: "",
			Deleted:    0,
		},
		Username: "admin",
		Password: "123456",
	}
	sql, err := builder.buildInsert(&u)
	if err != nil {
		t.Error(err)
	}

	expected := "INSERT INTO USER (ID,CREATE_TIME,CREATE_USER,UPDATE_TIME,UPDATE_USER,DELETED,USERNAME,PASSWORD) VALUES (?,?,?,?,?,?,?,?)"

	if sql != expected {
		t.Errorf("build insert error,\n expected: %s\n, got: %s\n", expected, sql)
	}
}

func TestBuilderUpdateById(t *testing.T) {
	var builder SqlBuilder
	u := User{
		BaseModel: BaseModel{
			ID:         1,
			CreateTime: time.Now(),
			CreateUser: "",
			UpdateTime: time.Now(),
			UpdateUser: "",
			Deleted:    0,
		},
		Username: "admin",
		Password: "123456",
	}
	sql, err := builder.buildUpdateById(&u)
	if err != nil {
		t.Error(err)
	}

	expected := "UPDATE USER SET UPDATE_TIME=?,UPDATE_USER=?,DELETED=?,USERNAME=?,PASSWORD=? WHERE ID=?"

	if sql != expected {
		t.Errorf("build insert error,\n expected: %s,\n got: %s\n", expected, sql)
	}
}

func TestBuilderGetById(t *testing.T) {
	var builder SqlBuilder
	u := User{
		BaseModel: BaseModel{
			ID:         1,
			CreateTime: time.Now(),
			CreateUser: "",
			UpdateTime: time.Now(),
			UpdateUser: "",
			Deleted:    0,
		},
		Username: "admin",
		Password: "123456",
	}
	sql, err := builder.buildGetById(&u)
	if err != nil {
		t.Error(err)
	}

	expected := "SELECT ID,CREATE_TIME,CREATE_USER,UPDATE_TIME,UPDATE_USER,DELETED,USERNAME,PASSWORD FROM USER WHERE ID=?"

	if sql != expected {
		t.Errorf("build insert error,\n expected: %s,\n got: %s\n", expected, sql)
	}
}
