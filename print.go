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
	InternalLog           StartLogEnvInput
	PrintLog              StartLogEnvInput
	FunctionLog           StartLogEnvInput
	InternalFunctionPrint StartLogEnvInput
	TimeFormat            string
}

type StartLogEnvInput struct {
	Development bool
	Stage       bool
	Production  bool
}

type Logger struct {
	InternalPrint         StartLogEnvInput
	DebugPrint            StartLogEnvInput
	FunctionPrint         StartLogEnvInput
	InternalFunctionPrint StartLogEnvInput
	TimeFormat            string
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

func checkDebugPrint() bool {
	switch inutil.Enviroment {
	case Enviroment_Development:
		if inutil.Logger.DebugPrint.Development {
			return true
		}
		return false
	case Enviroment_Stage:
		if inutil.Logger.DebugPrint.Stage {
			return true
		}
		return false
	case Enviroment_Production:
		if inutil.Logger.DebugPrint.Production {
			return true
		}
		return false
	}
	return false
}

func checkInternalPrint() bool {
	switch inutil.Enviroment {
	case Enviroment_Development:
		if inutil.Logger.InternalPrint.Development {
			return true
		}
		return false
	case Enviroment_Stage:
		if inutil.Logger.InternalPrint.Stage {
			return true
		}
		return false
	case Enviroment_Production:
		if inutil.Logger.InternalPrint.Production {
			return true
		}
		return false
	}
	return false
}

func checkFunctionPrint() bool {
	switch inutil.Enviroment {
	case Enviroment_Development:
		if inutil.Logger.InternalPrint.Development {
			return true
		}
		return false
	case Enviroment_Stage:
		if inutil.Logger.InternalPrint.Stage {
			return true
		}
		return false
	case Enviroment_Production:
		if inutil.Logger.InternalPrint.Production {
			return true
		}
		return false
	}
	return false
}

func logInternal(values ...any) {
	if checkInternalPrint() {
		Print(values...)
	}
}

func logInternalF(format string, values ...any) {
	if checkInternalPrint() {
		PrintF(format, values...)
	}
}

func LogDebug(values ...any) {
	if checkDebugPrint() {
		Print(values...)
	}
}

func LogDebugF(format string, values ...any) {
	if checkDebugPrint() {
		PrintF(format, values...)
	}
}

func SprintF(format string, values ...any) string {
	return fmt.Sprintf(format, values)
}

func Print(values ...any) {
	values = append(values, "\n")
	fmt.Print(values...)
}

func PrintF(format string, values ...any) {
	fmt.Printf(format+"\n", values...)
}

func PrintError(err error) {
	PrintErrorF("%v\n", err)
}

func PrintErrorF(format string, err error, values ...any) {
	values = append([]any{err.Error()}, values...)
	color.Red(logTime()+" "+format, values...)
}

func PrintFunction() func(string) {
	funcName := CallerName(2)
	if checkFunctionPrint() {
		PrintF("START: %v", funcName)
	}
	return func(funcName string) {
		if checkFunctionPrint() {
			PrintF("END: %v", funcName)
		}
	}
}

func PrintInternalFunction() func(string) {
	funcName := CallerName(2)
	if checkFunctionPrint() {
		PrintF("INTERNAL START: %v", funcName)
	}
	return func(funcName string) {
		if checkFunctionPrint() {
			PrintF("INTERNAL END: %v", funcName)
		}
	}
}
