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

type StartLogInput struct {
	InternalLog StartLogEnvInput
	DebugLog    StartLogEnvInput
	TimeFormat  string
}

type StartLogEnvInput struct {
	Development bool
	Stage       bool
	Production  bool
}

type Logger struct {
	InternalLog     StartLogEnvInput
	internalLogEnvs []string
	DebugLog        StartLogEnvInput
	debugEnvs       []string
	TimeFormat      string
}

type LogEnv struct {
	Development bool
	Stage       bool
	Production  bool
}

var clear map[string]func() //create a map for storing clear funcs

func startPrint() {
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

func setupDebug() {
	inutil.Logger.debugEnvs = []string{}
	if inutil.Logger.DebugLog.Development {
		inutil.Logger.debugEnvs = append(inutil.Logger.debugEnvs, Enviroment_Development)
	}
	if inutil.Logger.DebugLog.Stage {
		inutil.Logger.debugEnvs = append(inutil.Logger.debugEnvs, Enviroment_Stage)
	}
	if inutil.Logger.DebugLog.Production {
		inutil.Logger.debugEnvs = append(inutil.Logger.debugEnvs, Enviroment_Production)
	}
}

func setupInternalLog() {
	inutil.Logger.internalLogEnvs = []string{}
	if inutil.Logger.InternalLog.Development {
		inutil.Logger.internalLogEnvs = append(inutil.Logger.internalLogEnvs, Enviroment_Development)
	}
	if inutil.Logger.InternalLog.Stage {
		inutil.Logger.internalLogEnvs = append(inutil.Logger.internalLogEnvs, Enviroment_Stage)
	}
	if inutil.Logger.InternalLog.Production {
		inutil.Logger.internalLogEnvs = append(inutil.Logger.internalLogEnvs, Enviroment_Production)
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

func PrettyString(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func logTime() string {
	return time.Now().Format(inutil.Logger.TimeFormat)
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
	for _, env := range inutil.Logger.debugEnvs {
		if env == inutil.Enviroment {
			return true
		}
	}
	return false
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
	for _, env := range inutil.Logger.internalLogEnvs {
		if env == inutil.Enviroment {
			return true
		}
	}
	return false
}
