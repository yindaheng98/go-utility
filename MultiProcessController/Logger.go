package MultiProcessController

//A struct for logging.
type Logger interface {
	Log(...interface{})
}

type iLogger struct {
	Logger
	i uint64
}

func (log iLogger) Log(args ...interface{}) {
	log.Logger.Log(append([]interface{}{log.i, "-->"}, args...)...)
}
