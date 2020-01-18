package result

import "testing"

func TestSuccessJSON(t *testing.T) {
	inData := map[string]string{
		"name": "yangsen",
		"age":  "31",
	}
	expectOut := `{"code":0,"message":"成功","data":{"age":"31","name":"yangsen"}}`

	s := SuccessJSON(inData)
	if s != expectOut {
		t.Errorf("BuildSuccJSON failed!===>>>%s\n", s)
	}
}

func TestErrorJSON(t *testing.T) {
	inMsg := "失败"

	expectOut := `{"code":-1,"message":"失败"}`

	s := ErrorJSON(inMsg)
	if s != expectOut {
		t.Errorf("BuildErrorJSON failed!===>>>%s\n", s)
	}
}
