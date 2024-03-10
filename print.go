package inutil

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func logTime() string {
	return time.Now().Format(LogFormat)
}

func Log(values ...any) {
	LogF("", values...)
}

func LogF(format string, values ...any) {
	fmt.Printf(logTime()+format+"\n", values...)
}

func Error(err error) {
	ErrorF("", err)
}

func ErrorF(format string, err error) {
	color.Red(logTime()+"%v", err.Error())
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
	for _, env := range startModel.debugEnvs {
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
	for _, env := range startModel.internalLogEnvs {
		if env == startModel.Enviroment {
			return true
		}
	}
	return false
}
