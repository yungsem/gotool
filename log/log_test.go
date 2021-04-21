package log

import (
	"errors"
	"testing"
)

func TestLog_Info(t *testing.T) {
	log := NewLog("s", "debug", "d:\\project\\logs\\go-tool")
	err := errors.New("info test")
	log.Error(err)

	log.Info("test info")
	log.Info("test info: %s", "yangsen")

	log.Debug("test info")
	log.Debug("test info: %s", "yangsen")
}
