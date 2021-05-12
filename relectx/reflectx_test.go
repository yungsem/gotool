package reflectx

import (
	"fmt"
	"testing"
	"time"
)

type Base struct {
	Id int
	CreateTime time.Time
	UpdateTime time.Time
}

type User struct {
	Base
	Username string `cond:"like"`
	Password string `cond:"="`
	RealName string `cond:">="`
}

func TestGetFiledNamesSnake(t *testing.T) {
	u := &User{
		Username: "admin",
		Password: "123456",
		RealName: "管理员",
	}
	names, err := FiledNames(u)
	if err != nil {
		t.Errorf("failed: %s", err)
	}
	fmt.Println(names)
}

func TestGetStructName(t *testing.T) {
	u := &User{
		Username: "admin",
		Password: "123456",
		RealName: "管理员",
	}
	expected := "user"
	name, err := StructName(u)
	if err != nil {
		t.Error(err)
	}
	if name != expected {
		t.Errorf("expected: %s\ngot: %s\n", expected, name)
	}
}

func TestFiledNameTagMapSnakeLower(t *testing.T) {
	u := &User{
		Username: "admin",
		Password: "123456",
		RealName: "管理员",
	}
	m, err := FiledNameTagMapSnakeLower(u, "cond")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}