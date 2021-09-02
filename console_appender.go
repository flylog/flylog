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
	"io"
	"os"
	"runtime"
	"time"
)

func newConsoleAppender() appender {
	return new(consoleAppender)
}

type consoleAppender struct {
	*baseAppender
}

func (c *consoleAppender) WriteMessage(v string) {
	io.WriteString(os.Stderr, v)
}

func (b *consoleAppender) Layout(formart string, message ...interface{}) string {
	src := ""
	if GetLogLevel() < Info {
		pc, _, lineno, ok := runtime.Caller(3)
		if ok {
			src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
		}
	}
	msg := fmt.Sprint(message...)
	return fmt.Sprintf("[%s] -- %s -- %s", time.Now().Local().Format("2006-01-01 01:01:01"), src, msg)
}
