package flylog

type Logger interface {
	Trace(message ...interface{})
	Debug(message ...interface{})
	Info(message ...interface{})
	Warn(message ...interface{})
	Error(message ...interface{})
	Fatal(message ...interface{})
}

type LoggerFactory map[string]Logger

var loggerFactory = make(LoggerFactory)

func register(name string, l Logger) {
	loggerFactory[name] = l
}

func GetLogger() Logger {
	name := GlobalConf.logframeName
	return loggerFactory[name]
}
