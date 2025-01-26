package log

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
	return time.Now().Format(global.inutil.Logger.TimeFormat)
}

func checkDebugPrint() bool {
	switch global.inutil.Enviroment {
	case model.Enviroment_Development:
		if global.inutil.Logger.DebugPrint.Development {
			return true
		}
		return false
	case model.Enviroment_Stage:
		if global.inutil.Logger.DebugPrint.Stage {
			return true
		}
		return false
	case model.Enviroment_Production:
		if global.inutil.Logger.DebugPrint.Production {
			return true
		}
		return false
	}
	return false
}

func checkInternalLog() bool {
	switch global.inutil.Enviroment {
	case model.Enviroment_Development:
		if global.inutil.Logger.InternalPrint.Development {
			return true
		}
		return false
	case model.Enviroment_Stage:
		if global.inutil.Logger.InternalPrint.Stage {
			return true
		}
		return false
	case model.Enviroment_Production:
		if global.inutil.Logger.InternalPrint.Production {
			return true
		}
		return false
	}
	return false
}

func checkFunctionPrint() bool {
	switch global.inutil.Enviroment {
	case model.Enviroment_Development:
		if global.inutil.Logger.InternalPrint.Development {
			return true
		}
		return false
	case model.Enviroment_Stage:
		if global.inutil.Logger.InternalPrint.Stage {
			return true
		}
		return false
	case model.Enviroment_Production:
		if global.inutil.Logger.InternalPrint.Production {
			return true
		}
		return false
	}
	return false
}

func logInternal(values ...any) {
	if checkInternalLog() {
		Log(values...)
	}
}

func logInternalF(format string, values ...any) {
	if checkInternalLog() {
		LogF(format, values...)
	}
}

func LogDebug(values ...any) {
	if checkDebugPrint() {
		Log(values...)
	}
}

func LogDebugF(format string, values ...any) {
	if checkDebugPrint() {
		LogF(format, values...)
	}
}

func SprintF(format string, values ...any) string {
	return fmt.Sprintf(format, values)
}

func Log(values ...any) {
	format := ""
	for range values {
		format += "%v "
	}
	LogF(format, values...)
}

func LogF(format string, values ...any) {
	fmt.Print(logTime())
	fmt.Printf(format+"\n", values...)
}

func LogError(err Error) {
	LogErrorF("%v\n", err.Error())
}

func LogErrorF(format string, values ...any) {
	color.Red(logTime()+format, values...)
}

// func LogFunction() func() {
// 	funcName := CallerName(2)
// 	if checkFunctionPrint() {
// 		LogF("START: %v", funcName)
// 	}
// 	return func() {
// 		if checkFunctionPrint() {
// 			LogF("END: %v", funcName)
// 		}
// 	}
// }

// func LogInternalFunction() func() {
// 	funcName := CallerName(2)
// 	if checkFunctionPrint() {
// 		LogF("INTERNAL START: %v", funcName)
// 	}
// 	return func() {
// 		if checkFunctionPrint() {
// 			LogF("INTERNAL END: %v", funcName)
// 		}
// 	}
// }
