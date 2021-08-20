package main

import (
	"log"

	"github.com/flylog/flylog"
)

func main() {
	var i = 1
	var j = 2
	if j > i {

		log2 := flylog.GetLogger()
		log2.Debug("example flylog test")
		log.Println()
		log.Println("example flylog test2")

	}
}
