package inutil

type Return[V any] struct {
	Message string `json:"message"`
	Data    *V     `json:"data"`
	Success bool   `json:"success"`
	Status  int    `json:"-"`
}

type StartInput struct {
	Server     *StartServerInput
	Log        *StartLogInput
	Enviroment string
}

type StartLogInput struct {
	InternalLog     StartLogEnvInput
	internalLogEnvs []string
	DebugLog        StartLogEnvInput
	debugEnvs       []string
	TimeFormat      string
}

type StartLogEnvInput struct {
	Development bool
	Stage       bool
	Production  bool
}

type StartServerInput struct {
	Port string
}

type Inutil struct {
	Server *Server
}

var startModel *StartInput

func Start(start *StartInput) Inutil {
	out := Inutil{}

	Clear()

	startModel = start

	if startModel.Log != nil {
		setupDebug()
		setupInternalLog()
		if startModel.Log.TimeFormat == "" {
			startModel.Log.TimeFormat = LogFormat
			Log("No log time format specified, using default")
		}
	} else {
		startModel.Log = &StartLogInput{
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
		setupDebug()
		setupInternalLog()
		Log("No log specified, using default")
	}

	Log("Initializing...")

	startRequest()

	if startModel.Server != nil {
		logInternal("Starting server")
		out.Server = startModel.Server.start()
	}

	return out
}

func setupDebug() {
	startModel.Log.debugEnvs = []string{}
	if startModel.Log.DebugLog.Development {
		startModel.Log.debugEnvs = append(startModel.Log.debugEnvs, Enviroment_Development)
	}
	if startModel.Log.DebugLog.Stage {
		startModel.Log.debugEnvs = append(startModel.Log.debugEnvs, Enviroment_Stage)
	}
	if startModel.Log.DebugLog.Production {
		startModel.Log.debugEnvs = append(startModel.Log.debugEnvs, Enviroment_Production)
	}
}

func setupInternalLog() {
	startModel.Log.internalLogEnvs = []string{}
	if startModel.Log.InternalLog.Development {
		startModel.Log.internalLogEnvs = append(startModel.Log.internalLogEnvs, Enviroment_Development)
	}
	if startModel.Log.InternalLog.Stage {
		startModel.Log.internalLogEnvs = append(startModel.Log.internalLogEnvs, Enviroment_Stage)
	}
	if startModel.Log.InternalLog.Production {
		startModel.Log.internalLogEnvs = append(startModel.Log.internalLogEnvs, Enviroment_Production)
	}
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
