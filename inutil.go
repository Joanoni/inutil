package inutil

var startModel *Start_Model

func Start(start *Start_Model) Inutil {
	out := Inutil{}

	startModel = start
	startModel.debugEnvs = []string{}
	setupDebug()
	startModel.internalLogEnvs = []string{}
	setupInternalLog()

	if startModel.Server != nil {
		out.Server = server_Start(startModel.Server)
	}

	return out
}

func setupDebug() {
	if startModel.DebugLog.Development {
		startModel.debugEnvs = append(startModel.debugEnvs, Enviroment_Development)
	}
	if startModel.DebugLog.Stage {
		startModel.debugEnvs = append(startModel.debugEnvs, Enviroment_Stage)
	}
	if startModel.DebugLog.Production {
		startModel.debugEnvs = append(startModel.debugEnvs, Enviroment_Production)
	}
}

func setupInternalLog() {
	if startModel.InternalLog.Development {
		startModel.internalLogEnvs = append(startModel.internalLogEnvs, Enviroment_Development)
	}
	if startModel.InternalLog.Stage {
		startModel.internalLogEnvs = append(startModel.internalLogEnvs, Enviroment_Stage)
	}
	if startModel.InternalLog.Production {
		startModel.internalLogEnvs = append(startModel.internalLogEnvs, Enviroment_Production)
	}
}
