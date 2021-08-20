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
)

func newFileAppender(fileName string) appender {
	fileAppend := new(fileAppender)
	fileAppend.fileName = "logs/" + fileName
	return fileAppend
}

type fileAppender struct {
	*baseAppender
	fileName string
}

func (f *fileAppender) WriteMessage(v string) {
	_, err := os.Stat(f.fileName)
	if os.IsNotExist(err) {
		os.Mkdir("logs/", os.ModePerm)
	}

	file, err := os.Create(f.fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("os.Create %s failed: %v", file, err)
		return
	}
	io.WriteString(file, fmt.Sprintf("%s", v))
}
