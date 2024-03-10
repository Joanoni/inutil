package inutil

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Clear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

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
		LogF(format, values...)
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
		LogF(format, values...)
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
