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
	"log"

	"github.com/obity/properties"
)

func ToLevel(s string) Level {
	switch s {
	case "Trace", "trace":
		return Trace
	case "Debug", "debug":
		return Debug
	case "Info", "info":
		return Info
	case "Warn", "warn":
		return Warn
	case "Error", "error":
		return Error
	case "Fatal", "fatal":
		return Fatal
	default:
		return Info
	}
}

type Conf struct {
	logframeName    string
	logLevel        Level
	confFile        string
	logPrefix       string
	formart         string
	ConsoleAppender string
	FileAppender    string
	Appenders       []Appender
}

func init() {
	var (
		isExist bool
	)
	GlobalConf.confFile = ConfFile
	p := properties.NewProperties()
	p.LoadFromFile(ConfFile)

	level, isExist := p.Property("LogLevel")
	GlobalConf.logLevel = ToLevel(level)
	if !isExist {
		log.Println("LogLevel is not exist")
	}
	GlobalConf.logPrefix, isExist = p.Property("LogPrefix")
	if !isExist {
		log.Println("LogPrefix is not exist")
	}

	GlobalConf.formart, isExist = p.Property("Formart")
	if !isExist {
		// log.Println("Formart is not exist")
	}
	GlobalConf.ConsoleAppender, isExist = p.Property("ConsoleAppender")
	if !isExist {
		log.Println("ConsoleAppender is not exist")
	}
	GlobalConf.logframeName, isExist = p.Property("LogframeName")
	if !isExist {
		log.Println("LogframeName is not exist")
	}

	GlobalConf.FileAppender, isExist = p.Property("FileAppender")
	if !isExist {
		log.Println("FileAppender is not exist")
	}
}
