package flylog

var GlobalConf Conf

type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

const ConfFile = "flylog.properties"

func (l Level) String() string {
	switch l {
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Info"
	}
}
