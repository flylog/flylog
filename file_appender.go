package flylog

import (
	"fmt"
	"io"
	"os"
)

func NewFileAppender(fileName string) Appender {
	fileAppender := new(FileAppender)
	fileAppender.fileName = "logs/" + fileName
	return fileAppender
}

type FileAppender struct {
	*BaseAppender
	fileName string
}

func (f *FileAppender) WriteMessage(v string) {
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
