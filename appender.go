package flylog

import (
	"fmt"
	"runtime"
	"time"
)

type Appender interface {
	WriteMessage(v string)
	Filter(level Level) bool
	Layout(format string, message ...interface{}) string
}

type BaseAppender struct {
}

func (b *BaseAppender) Layout(formart string, message ...interface{}) string {
	src := ""
	if GetLogLevel() < Info {
		pc, _, lineno, ok := runtime.Caller(3)
		if ok {
			src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
		}
	}
	msg := fmt.Sprint(message...)
	return fmt.Sprintf("[%s] -- %s / %s", time.Now().Local().String(), src, msg)
}

func (b *BaseAppender) Filter(level Level) bool {
	if level >= GlobalConf.logLevel {
		return true
	}
	return false
}
