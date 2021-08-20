package main

import (
	"github.com/flylog/flylog"
)

func main() {

	logs := flylog.GetLogger()
	for i := 0; i < 100; i++ {
		logs.Debug("example flylog test", i)
		logs.Info("example lylog test2", i)
		logs.Error("example lylog test2", i)
		logs.Warn("example lylog test2", i)
	}
}
