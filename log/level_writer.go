package log

import (
	"github.com/rs/zerolog"
	"github.com/yungsem/gotool/times"
	"os"
	"time"
)

type levelWriterImpl struct {
	logPath string
}

func (lw levelWriterImpl) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (lw levelWriterImpl) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	// 文件名
	var fileName string
	switch level {
	case zerolog.ErrorLevel:
		fileName = "error.log"
	case zerolog.InfoLevel:
		fileName = "info.log"
	case zerolog.DebugLevel:
		fileName = "debug.log"
	case zerolog.WarnLevel:
		fileName = "warn.log"
	}

	// 获取当前日期
	now := time.Now()
	today := times.Format(now, times.LayoutDate)
	todayTight := times.Format(now, times.LayoutDateTight)
	// 目录分割符
	sep := string(os.PathSeparator)
	// 目录路径
	dirPath := lw.logPath + sep + today + sep
	// 文件路径
	filePath := dirPath + todayTight + "_" + fileName

	// 创建目录
	// 如果目录已经存在，则什么都不做
	err = os.MkdirAll(dirPath, os.ModeDir)
	if err != nil {
		return
	}

	// 创建文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	// 写入
	return file.Write(p)
}
