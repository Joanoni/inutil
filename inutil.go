package inutil

var startModel *Start_Model

func Start(start *Start_Model) Inutil {
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
		startModel.Log = &Start_Log{
			InternalLog: Start_Log_Envs{
				Development: true,
				Stage:       false,
				Production:  false,
			},
			DebugLog: Start_Log_Envs{
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

	if startModel.Server != nil {
		internalLog("Starting server")
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
