package inutil

import "log"

type Return[V any] struct {
	Message    string `json:"message"`
	Data       *V     `json:"data"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"-"`
}

type StartInput struct {
	Server     *StartServerInput
	Log        *StartLogInput
	WebSocket  *StartWebSocketInput
	Enviroment string
}

type Inutil struct {
	Server           *Server
	Logger           *Logger
	WebSocketManager *WebSocketManager
	Enviroment       string
}

var inutil Inutil

func init() {
	log.Println("Initializing...")
	startPrint()
	startRequest()
}

func Start(start *StartInput) Inutil {
	inutil = Inutil{
		Enviroment: start.Enviroment,
	}

	Clear()

	if start.Log == nil {
		start.Log = &StartLogInput{
			InternalLog: StartLogEnvInput{
				Development: true,
				Stage:       false,
				Production:  false,
			},
			DebugLog: StartLogEnvInput{
				Development: true,
				Stage:       true,
				Production:  false,
			},
			TimeFormat: LogFormat,
		}
		Log("No log specified, using default")
	} else {
		if start.Log.TimeFormat == "" {
			start.Log.TimeFormat = LogFormat
			Log("No log time format specified, using default")
		}
	}

	inutil.Logger = &Logger{
		InternalLog: start.Log.InternalLog,
		DebugLog:    start.Log.DebugLog,
		TimeFormat:  start.Log.TimeFormat,
	}
	setupLogger()
	setupInternalLog()

	if start.Server != nil {
		logInternal("Starting Server")
		inutil.Server = start.Server.start()
	}

	if start.WebSocket != nil {
		logInternal("Starting WebSocketManager")
		inutil.WebSocketManager = start.WebSocket.startWebSocket()
	}

	return inutil
}

func JSON[T any](payload Return[T], c *Context) {
	c.gc.JSON(payload.StatusCode, payload)
}

const (
	Enviroment_Development = "development"
	Enviroment_Stage       = "stage"
	Enviroment_Production  = "production"

	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
	LogFormat  = "2006-01-02 15:04:05.000"
)
