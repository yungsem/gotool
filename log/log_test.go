package log

import "testing"

func TestLog_Info(t *testing.T) {
	log := NewLog("stdio", "debug", "C:\\ys\\project\\go_space\\logs", "gotool")
	log.Info("测试日志")
}

func TestLog_InfoF(t *testing.T) {
	log := NewLog("stdio", "debug", "C:\\ys\\project\\go_space\\logs", "gotool")
	log.InfoF("测试%s日志%s", "gotool", "yungsem")
}