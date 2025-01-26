package inutil

import "log"

type StartInput struct {
	Server     *model.StartServerInput
	Log        *model.StartLogInput
	WebSocket  *model.StartWebSocketInput
	Enviroment string
}

func init() {
	log.Println("Initializing...")
	startPrint()
	startRequest()
}

func Start(start *StartInput) model.Inutil {
	inutil = model.Inutil{
		Enviroment: start.Enviroment,
	}

	inprint.Clear()

	if start.Log == nil {
		start.Log = &model.StartLogInput{
			InternalLog: model.StartLogEnvInput{
				Development: true,
				Stage:       false,
				Production:  false,
			},
			PrintLog: model.StartLogEnvInput{
				Development: true,
				Stage:       true,
				Production:  false,
			},
			TimeFormat: LogFormat,
		}
		Print("No log specified, using default")
	} else {
		if start.Log.TimeFormat == "" {
			start.Log.TimeFormat = LogFormat
			Print("No log time format specified, using default")
		}
	}

	inutil.Logger = &model.Logger{
		InternalPrint:         start.Log.InternalLog,
		DebugPrint:            start.Log.PrintLog,
		FunctionPrint:         start.Log.FunctionLog,
		InternalFunctionPrint: start.Log.InternalFunctionPrint,
		TimeFormat:            start.Log.TimeFormat,
	}
	// setupPrint()
	// setupInternalPrint()

	if start.Server != nil {
		logInternal("Starting Server")
		inutil.Server = start.Server.start()
		inutil.Server.Use(MiddlewareRecovery())
	}

	if start.WebSocket != nil {
		logInternal("Starting WebSocketManager")
		inutil.WebSocketManager = start.WebSocket.startWebSocket()
	}

	return inutil
}
