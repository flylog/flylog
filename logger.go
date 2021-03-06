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

type Logger interface {
	Trace(message ...interface{})
	Debug(message ...interface{})
	Info(message ...interface{})
	Warn(message ...interface{})
	Error(message ...interface{})
	Fatal(message ...interface{})
}

type factory map[string]Logger

var globalFactory = make(factory)

func register(name string, l Logger) {
	globalFactory[name] = l
}

func GetLogger() Logger {
	name := globalFlyConf.logframeName
	return globalFactory[name]
}
