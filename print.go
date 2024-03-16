package inutil

import (
	"encoding/json"
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

func LogDebugPretty(v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	LogDebug(string(b))
	return nil
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

func LogError(err error) {
	LogErrorF("%v", err)
}

func LogErrorF(format string, err error, values ...any) {
	values = append([]any{err.Error()}, values...)
	color.Red(logTime()+" "+format, values...)
}

func LogDebug(values ...any) {
	if checkDebug() {
		Log(values...)
	}
}

func LogDebugF(format string, values ...any) {
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

func logInternalPretty(v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	logInternal(string(b))
	return nil
}

func logInternal(values ...any) {
	if checkLogInternal() {
		Log(values...)
	}
}

func logInternalF(format string, values ...any) {
	if checkLogInternal() {
		LogF(format, values...)
	}
}

func checkLogInternal() bool {
	for _, env := range startModel.Log.internalLogEnvs {
		if env == startModel.Enviroment {
			return true
		}
	}
	return false
}
