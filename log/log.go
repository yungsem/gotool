package log

import (
	"github.com/rs/zerolog"
	"github.com/yungsem/gotool/times"
	"os"
)

const OutputFile = "file"

type Log struct {
	zeroLog    zerolog.Logger
	serverName string
}

func (l *Log) Info(msg string) {
	l.zeroLog.Info().Msg(msg)
}

func (l *Log) InfoF(format string, v ...interface{}) {
	l.zeroLog.Info().Msgf(format, v...)
}

func (l *Log) Debug(msg string) {
	l.zeroLog.Debug().Msg(msg)
}

func (l *Log) DebugF(format string, v ...interface{}) {
	l.zeroLog.Debug().Msgf(format, v...)
}

func (l *Log) Error(err error) {
	l.zeroLog.Error().Err(err).Msg(err.Error())
}

func NewLog(output string, level string, logPath string, serverName string) *Log {
	var zeroLog zerolog.Logger

	// 设置日志的输出路径
	if output == OutputFile { // 输出到文件
		w := levelWriterImpl{
			logPath:    logPath,
			serverName: serverName,
		}
		zeroLog = zerolog.New(w).With().Timestamp().Logger()
	} else { // 输出到控制台
		zeroLog = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	// 设置日志的时间格式
	zerolog.TimeFieldFormat = times.LayoutStandardMil

	// 设置日志级别
	switch level {
	case zerolog.DebugLevel.String():
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case zerolog.InfoLevel.String():
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// 构建 log
	log := &Log{
		zeroLog:    zeroLog,
		serverName: serverName,
	}

	return log
}
