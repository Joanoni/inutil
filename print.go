package inutil

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func logTime() string {
	return time.Now().Format(startModel.Log.TimeFormat)
}

func Log(values ...any) {
	format := ""
	for i := 0; i < len(values); i++ {
		format += "%v"
	}
	LogF(format, values...)
}

func LogF(format string, values ...any) {
	fmt.Printf(logTime()+" "+format+"\n", values...)
}

func Error(err error) {
	ErrorF("%v", err)
}

func ErrorF(format string, err error, values ...any) {
	values = append([]any{err.Error()}, values...)
	color.Red(logTime()+" "+format, values...)
}

func Debug(values ...any) {
	if checkDebug() {
		Log(values...)
	}
}

func DebugF(format string, values ...any) {
	if checkDebug() {
		LogF(format+"\n", values...)
	}
}

func checkDebug() bool {
	for _, env := range startModel.Log.debugEnvs {
		if env == startModel.Enviroment {
			return true
		}
	}
	return false
}

func internalLog(values ...any) {
	if checkInternalLog() {
		Log(values...)
	}
}

func internalLogF(format string, values ...any) {
	if checkInternalLog() {
		LogF(format+"\n", values...)
	}
}

func checkInternalLog() bool {
	for _, env := range startModel.Log.internalLogEnvs {
		if env == startModel.Enviroment {
			return true
		}
	}
	return false
}
