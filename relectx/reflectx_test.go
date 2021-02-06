package reflectx

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
	Password string
	RealName string
}

func TestGetFiledNamesSnake(t *testing.T) {
	u := &User{
		Username: "admin",
		Password: "123456",
		RealName: "管理员",
	}
	names, err := GetFiledNamesSnakeLower(u)
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
	name, err := GetStructName(u)
	if err != nil {
		t.Error(err)
	}
	if name != expected {
		t.Errorf("expected: %s\ngot: %s\n", expected, name)
	}
}
