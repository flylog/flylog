package flylog

import (
	"io"
	"os"
)

func NewConsoleAppender() Appender {
	return new(ConsoleAppender)
}

type ConsoleAppender struct {
	*BaseAppender
}

func (c *ConsoleAppender) WriteMessage(v string) {
	io.WriteString(os.Stderr, v)
}
