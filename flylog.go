package flylog

func SetLevel(level Level) {
	GlobalConf.logLevel = level

}

func GetLogLevel() Level {
	return GlobalConf.logLevel
}

func SetPrefix(prefix string) {
	GlobalConf.logPrefix = prefix
}

func Prefix() string {
	return GlobalConf.logPrefix
}

func Formart() string {
	return GlobalConf.formart
}
