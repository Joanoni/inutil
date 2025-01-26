package model

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
