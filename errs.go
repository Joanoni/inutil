package inutil

type Errs struct {
	Message string
}

func HandleError(err error) bool {
	if err != nil {
		PrintErrsF("HandleError: %v", err)
		return true
	}
	return false
}

func HandleErrs(err *Errs) bool {
	if err != nil {
		PrintErrsF("HandleError: %v", err.Message)
		return true
	}
	return false
}

func ErrsFromError(err error) Errs {
	return Errs{Message: err.Error()}
}

var (
	ErrsContentTypeNotSet = Errs{Message: "Content-Type header not set"}
)
