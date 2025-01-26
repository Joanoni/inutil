package inutil

import "runtime"

func CallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		Print("error: CallerName")
		return ""
	}
	return runtime.FuncForPC(pc).Name()
}
