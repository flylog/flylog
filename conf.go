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

type flyConf struct {
	logframeName    string
	logLevel        Level
	confFile        string
	logPrefix       string
	formart         string
	ConsoleAppender string
	FileAppender    string
	Appenders       []appender
}

func init() {
	var (
		isExist bool
	)
	globalFlyConf.confFile = flyPropertyFile
	p := properties.NewProperties()
	p.LoadFromFile(flyPropertyFile)

	level, isExist := p.Property("LogLevel")
	globalFlyConf.logLevel = ToLevel(level)
	if !isExist {
		log.Println("LogLevel is not exist")
	}
	globalFlyConf.logPrefix, isExist = p.Property("LogPrefix")
	if !isExist {
		log.Println("LogPrefix is not exist")
	}

	globalFlyConf.formart, isExist = p.Property("Formart")
	if !isExist {
		// log.Println("Formart is not exist")
	}
	globalFlyConf.ConsoleAppender, isExist = p.Property("ConsoleAppender")
	if !isExist {
		log.Println("ConsoleAppender is not exist")
	}
	globalFlyConf.logframeName, isExist = p.Property("LogframeName")
	if !isExist {
		log.Println("LogframeName is not exist")
	}

	globalFlyConf.FileAppender, isExist = p.Property("FileAppender")
	if !isExist {
		log.Println("FileAppender is not exist")
	}
}
