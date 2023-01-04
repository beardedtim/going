package log

import (
	"fmt"
	"os"
)

type ILogger interface {
	Log()
	Info()
	Error()
	Warn()
	Print()
}

type Logger struct {
	ILogger
	prefix string
	level  int
}

func mapLevelToName(level int) string {
	switch level {
	case 0:
		return "LOG"
	case 10:
		return "INFO"
	case 20:
		return "DEBUG"
	case 30:
		return "WARN"
	case 40:
		return "ERROR"
	default:
		return "DEBUG"
	}
}

func (l Logger) Print(level int, msg string) {
	if l.level <= level {
		fmt.Println(fmt.Sprintf("[%d][%s] [%s]: %s", os.Getpid(), l.prefix, mapLevelToName(level), msg))
	}
}

func (l Logger) Log(msg string) {
	level := 0

	l.Print(level, msg)
}

func (l Logger) Info(msg string) {
	level := 10

	l.Print(level, msg)
}

func (l Logger) Debug(msg string) {
	level := 20

	l.Print(level, msg)
}

func (l Logger) Warn(msg string) {
	level := 30

	l.Print(level, msg)
}

func (l Logger) Error(msg string) {
	level := 40

	l.Print(level, msg)
}

func CreateLogger(prefix string) Logger {
	log := Logger{prefix: prefix, level: 0}

	return log
}
