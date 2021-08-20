package flylog

import (
	"fmt"
	"time"
)

func init() {
	logFrame := newLog4g()
	if logFrame.Name() == GlobalConf.logframeName {
		if len(GlobalConf.ConsoleAppender) > 0 {
			logFrame.outs = append(logFrame.outs, NewConsoleAppender())
		}
		if len(GlobalConf.FileAppender) > 0 {
			now := time.Now().Format("2006-02-05_")
			logFrame.outs = append(logFrame.outs, NewFileAppender(now+GlobalConf.FileAppender))
		}
	}
	register(logFrame.Name(), logFrame)
}

func newLog4g() *Log4g {
	l := new(Log4g)
	return l
}

type Log4g struct {
	outs []Appender
}

func (l *Log4g) Name() string {
	return "Log4g"
}

func (l *Log4g) Trace(message ...interface{}) {
	l.log(Trace, message...)
}

func (l *Log4g) Debug(message ...interface{}) {
	l.log(Debug, message...)
}

func (l *Log4g) Info(message ...interface{}) {
	l.log(Info, message...)
}

func (l *Log4g) Warn(message ...interface{}) {
	l.log(Warn, message...)
}

func (l *Log4g) Error(message ...interface{}) {
	l.log(Error, message...)
}

func (l *Log4g) Fatal(message ...interface{}) {
	l.log(Debug, message...)
}

func (l *Log4g) log(level Level, message ...interface{}) {
	for _, out := range l.outs {
		if !out.Filter(level) {
			continue
		}
		layoutStr := out.Layout(
			Formart(),
			fmt.Sprint(message...),
			fmt.Sprintf(" [%s] <%s> ", level.String(), Prefix()))
		out.WriteMessage(fmt.Sprintln(layoutStr))
	}
}
