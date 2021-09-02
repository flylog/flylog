/*
   Copyright (c) 2021 ffactory.org
   Flylog is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package flylog

import (
	"fmt"
	"time"

	"github.com/flylog/colorstyle"
)

func init() {
	logFrame := newLog4g()
	if logFrame.Name() == globalFlyConf.logframeName {
		if len(globalFlyConf.ConsoleAppender) > 0 {
			logFrame.outs = append(logFrame.outs, newConsoleAppender())
		}
		if len(globalFlyConf.FileAppender) > 0 {
			now := time.Now().Format("2006-01-01")
			fileName := now + "_" + globalFlyConf.FileAppender
			logFrame.outs = append(logFrame.outs, newFileAppender(fileName))
		}
	}
	register(logFrame.Name(), logFrame)
}

func newLog4g() *Log4g {
	l := new(Log4g)
	return l
}

type Log4g struct {
	outs []appender
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
			fmt.Sprintf(" [%s] <%s> ", colorstyle.Blue(level.String()), Prefix()))
		out.WriteMessage(fmt.Sprintln(layoutStr))

		/* if level == Error {
			panic(layoutStr)
		}
			 if level == Fatal {
		os.Exit(1)
		} */
	}
}
