package logger

import (
	"flag"
	"log"
	"os"
)

// Log for log to file
var Log *log.Logger


func init() {
	// set location of log file
	var logpath = "tmp/app.log"

	flag.Parse()
	var file, err = os.Create(logpath)

	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
