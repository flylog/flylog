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
	"runtime"
	"time"
)

type appender interface {
	WriteMessage(v string)
	Filter(level Level) bool
	Layout(format string, message ...interface{}) string
}

type baseAppender struct {
}

func (b *baseAppender) Layout(formart string, message ...interface{}) string {
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

func (b *baseAppender) Filter(level Level) bool {
	if level >= globalFlyConf.logLevel {
		return true
	}
	return false
}
