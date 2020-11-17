package util

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	const flags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix

	if "development" == os.Getenv("target") {
		Trace = log.New(os.Stdout, "TRACE: ", flags)
	} else {
		Trace = log.New(ioutil.Discard, "TRACE: ", flags)
	}

	Info = log.New(os.Stdout, "INFO: ", flags)
	Warning = log.New(os.Stdout, "WARNING: ", flags)
	Error = log.New(os.Stderr, "ERROR: ", flags)
}
