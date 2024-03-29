package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var entry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func CreateLogger() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = getFormatter()

	if err := os.MkdirAll("logs", 0o700); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}

	logger.SetOutput(io.Discard)
	logger.AddHook(&writerHook{
		Writer:    []io.Writer{logFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)
	logger.Infof("[OK] logger created: %+v", logger)

	entry = logrus.NewEntry(logger)
}

func CreateTestLogger() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = getFormatter()

	logger.SetOutput(io.Discard)
	logger.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetLevel(logrus.InfoLevel)
	logger.Infof("[OK] test logger created: %+v", logger)

	entry = logrus.NewEntry(logger)
}

func getFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return frame.Function, fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors:    true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
}

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return nil
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}
